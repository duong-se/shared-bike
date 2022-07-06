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

func (u *useCaseImpl) GetAllBike(ctx context.Context) ([]domain.GetAllBikeResponse, error) {
	u.logger.Info("[BikeUseCase.GetAllBike] fetching all bikes")
	bikes, err := u.repository.GetList(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.GetAllBikeResponse{}, nil
	}
	if err != nil {
		u.logger.Error("[BikeUseCase.GetAllBike] fetch all bikes failed", err)
		return []domain.GetAllBikeResponse{}, apperrors.ErrInternalServerError
	}
	userIDs := u.getUserIDs(bikes)
	usersMap, err := u.fetchMapUsersByID(ctx, userIDs)
	if err != nil {
		u.logger.Error("[BikeUseCase.GetAllBike] fetch user map failed", err)
		return []domain.GetAllBikeResponse{}, apperrors.ErrInternalServerError
	}
	result := u.transformGetAllBikeResponse(bikes, usersMap)
	u.logger.Info("[BikeUseCase.GetAllBike] fetch all bikes success")
	return result, nil
}

func (u *useCaseImpl) transformGetAllBikeResponse(bikes *[]domain.Bike, usersMap map[int64]domain.User) []domain.GetAllBikeResponse {
	results := []domain.GetAllBikeResponse{}
	for _, bike := range *bikes {
		renter := domain.GetAllBikeResponse{
			ID:     bike.ID,
			Lat:    bike.Lat.String(),
			Long:   bike.Long.String(),
			Status: bike.Status,
			UserID: bike.UserID,
		}
		if bike.UserID != nil {
			name := usersMap[*bike.UserID].Name
			username := usersMap[*bike.UserID].Username
			renter.NameOfRenter = &name
			renter.UsernameOfRenter = &username
		}
		results = append(results, renter)
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

func (u *useCaseImpl) Rent(ctx context.Context, payload domain.RentOrReturnRequestPayload) (*domain.Bike, error) {
	u.logger.Infof("[BikeUseCase.Rent] user %d is renting bike %d", payload.UserID, payload.ID)
	isRented, err := u.checkRented(ctx, payload.UserID)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] user %d check rented or not failed", payload.UserID, err)
		return nil, apperrors.ErrInternalServerError
	}
	if isRented {
		u.logger.Infof("[BikeUseCase.Rent] user %d is already renting a bike", payload.UserID)
		return nil, apperrors.ErrUserHaveBikeAlready
	}
	currentBike, err := u.repository.GetByID(ctx, payload.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Infof("[BikeUseCase.Rent] cannot find bike %d", payload.ID)
		return nil, apperrors.ErrBikeNotFound
	}
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] fetch current bike %d failed", payload.ID, err)
		return nil, apperrors.ErrInternalServerError
	}
	if currentBike.IsRented() {
		u.logger.Info("[BikeUseCase.Rent] cannot rent because bike is rented")
		return nil, apperrors.ErrBikeRented
	}
	updatedBike := &domain.Bike{
		ID:     currentBike.ID,
		Lat:    currentBike.Lat,
		Long:   currentBike.Long,
		Status: domain.BikeStatusRented,
		UserID: &payload.UserID,
	}
	err = u.repository.Update(ctx, updatedBike)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] user %d rent bike %d failed", payload.UserID, payload.ID, err)
		return nil, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Rent] user %d rent bike %d success", payload.UserID, payload.ID)
	return updatedBike, nil
}

func (u *useCaseImpl) Return(ctx context.Context, payload domain.RentOrReturnRequestPayload) (*domain.Bike, error) {
	currentBike, err := u.repository.GetByID(ctx, payload.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Infof("[BikeUseCase.Return] cannot find bike %d", payload.ID)
		return nil, apperrors.ErrBikeNotFound
	}
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Return] fetch current bike %d failed", payload.ID, err)
		return nil, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Return] user %d is returning bike %d", currentBike.UserID, payload.ID)
	if currentBike.IsAvailable() {
		u.logger.Info("[BikeUseCase.Return] cannot return because bike is available")
		return nil, apperrors.ErrBikeAvailable
	}
	if payload.UserID != *currentBike.UserID {
		u.logger.Info("[BikeUseCase.Return] cannot return because bike is not yours")
		return nil, apperrors.ErrBikeNotYours
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
		u.logger.Errorf("[BikeUseCase.Return] user %d is return bike %d failed", currentBike.UserID, payload.ID, err)
		return nil, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Return] user %d is return bike %d failed", currentBike.UserID, payload.ID)
	return updatedBike, nil
}
