package bike

import (
	"context"
	"errors"

	"shared-bike/apperrors"
	"shared-bike/domain"

	"gorm.io/gorm"
)

type useCaseImpl struct {
	repository     IRepository
	logger         ILogger
	userRepository IUserRepository
}

func NewUseCase(logger ILogger, repository IRepository, userRepository IUserRepository) *useCaseImpl {
	return &useCaseImpl{
		repository:     repository,
		logger:         logger,
		userRepository: userRepository,
	}
}

func (u *useCaseImpl) GetAllBike(ctx context.Context) ([]domain.BikeDTO, error) {
	u.logger.Info("[BikeUseCase.GetAllBike] fetching all bikes")
	bikes, err := u.repository.GetList(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.BikeDTO{}, nil
	}
	if err != nil {
		u.logger.Error("[BikeUseCase.GetAllBike] fetch all bikes failed", err)
		return []domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	userIDs := u.getUserIDs(bikes)
	usersMap, err := u.fetchMapUsersByID(ctx, userIDs)
	if err != nil {
		u.logger.Error("[BikeUseCase.GetAllBike] fetch user map failed", err)
		return []domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	result := u.transformBikeDTOList(bikes, usersMap)
	u.logger.Info("[BikeUseCase.GetAllBike] fetch all bikes success")
	return result, nil
}

func (u *useCaseImpl) transformBikeDTOList(bikes *[]domain.Bike, usersMap map[int64]domain.User) []domain.BikeDTO {
	results := []domain.BikeDTO{}
	for _, bike := range *bikes {
		rentedBike := bike.ToDTO()
		if (bike.UserID != nil && usersMap[*bike.UserID] != domain.User{}) {
			name := usersMap[*bike.UserID].Name
			username := usersMap[*bike.UserID].Username
			rentedBike.NameOfRenter = &name
			rentedBike.UsernameOfRenter = &username
		}
		results = append(results, rentedBike)
	}
	return results
}

func (u *useCaseImpl) fetchMapUsersByID(ctx context.Context, userIDs []int64) (map[int64]domain.User, error) {
	u.logger.Infof("[BikeUseCase.fetchUsers] fetch all user by IDs failed", userIDs)
	users, err := u.userRepository.GetListByIDs(ctx, userIDs)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return map[int64]domain.User{}, nil
	}
	if err != nil {
		u.logger.Error("[BikeUseCase.fetchUsers] fetch all bikes failed", err)
		return map[int64]domain.User{}, err
	}
	usersMap := map[int64]domain.User{}
	for _, user := range *users {
		usersMap[user.ID] = user
	}
	u.logger.Info("[BikeUseCase.fetchUsers] fetch all users success")
	return usersMap, nil
}

func (u *useCaseImpl) getUserIDs(bikes *[]domain.Bike) []int64 {
	var (
		userIDs []int64
	)
	for _, bike := range *bikes {
		if bike.UserID != nil {
			userIDs = append(userIDs, *bike.UserID)
		}
	}
	return userIDs
}

func (u *useCaseImpl) checkRented(ctx context.Context, userID int64) (bool, error) {
	total, err := u.repository.CountByUserID(ctx, userID)
	if err != nil {
		return true, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

func (u *useCaseImpl) Rent(ctx context.Context, body domain.RentOrReturnRequestPayload) (domain.BikeDTO, error) {
	u.logger.Infof("[BikeUseCase.Rent] user %d is renting bike %d", body.UserID, body.ID)
	isRented, err := u.checkRented(ctx, body.UserID)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] user %d check rented or not failed", body.UserID, err)
		return domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	if isRented {
		u.logger.Infof("[BikeUseCase.Rent] user %d is already renting a bike", body.UserID)
		return domain.BikeDTO{}, apperrors.ErrUserHasBikeAlready
	}
	currentUser, err := u.userRepository.GetByID(ctx, body.UserID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Errorf("[BikeUseCase.Rent] user %d not exists", body.UserID, err)
		return domain.BikeDTO{}, apperrors.ErrUserNotExisted
	}
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] user %d fetch failed", body.UserID, err)
		return domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	currentBike, err := u.repository.GetByID(ctx, body.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Infof("[BikeUseCase.Rent] cannot find bike %d", body.ID)
		return domain.BikeDTO{}, apperrors.ErrBikeNotFound
	}
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] fetch current bike %d failed", body.ID, err)
		return domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	if currentBike.IsRented() {
		u.logger.Info("[BikeUseCase.Rent] cannot rent because bike is rented")
		return domain.BikeDTO{}, apperrors.ErrBikeRented
	}
	updatedBike := &domain.Bike{
		ID:     currentBike.ID,
		Lat:    currentBike.Lat,
		Long:   currentBike.Long,
		Status: domain.BikeStatusRented,
		UserID: &body.UserID,
	}
	err = u.repository.Update(ctx, updatedBike)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] user %d rent bike %d failed", body.UserID, body.ID, err)
		return domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Rent] user %d rent bike %d success", body.UserID, body.ID)
	result := updatedBike.ToDTO()
	if currentUser != nil {
		result.NameOfRenter = &currentUser.Name
		result.UsernameOfRenter = &currentUser.Username
	}
	return result, nil
}

func (u *useCaseImpl) Return(ctx context.Context, body domain.RentOrReturnRequestPayload) (domain.BikeDTO, error) {
	currentBike, err := u.repository.GetByID(ctx, body.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Infof("[BikeUseCase.Return] cannot find bike %d", body.ID)
		return domain.BikeDTO{}, apperrors.ErrBikeNotFound
	}
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Return] fetch current bike %d failed", body.ID, err)
		return domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Return] user %d is returning bike %d", currentBike.UserID, body.ID)
	if currentBike.IsAvailable() {
		u.logger.Info("[BikeUseCase.Return] cannot return because bike is available")
		return domain.BikeDTO{}, apperrors.ErrBikeAvailable
	}
	if body.UserID != *currentBike.UserID {
		u.logger.Info("[BikeUseCase.Return] cannot return because bike is not yours")
		return domain.BikeDTO{}, apperrors.ErrBikeNotYours
	}

	updatedBike := &domain.Bike{
		ID:     currentBike.ID,
		Lat:    currentBike.Lat,
		Long:   currentBike.Long,
		Status: domain.BikeStatusAvailable,
		UserID: nil,
	}
	err = u.repository.Update(ctx, updatedBike)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Return] user %d is return bike %d failed", currentBike.UserID, body.ID, err)
		return domain.BikeDTO{}, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Return] user %d is return bike %d failed", currentBike.UserID, body.ID)
	result := updatedBike.ToDTO()
	return result, nil
}
