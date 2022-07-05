package bike

import (
	"context"
	"errors"

	"github.com/duong-se/shared-bike/apperrors"
	"github.com/duong-se/shared-bike/domain"
	"gorm.io/gorm"
)

type useCaseImpl struct {
	repo   IRepository
	logger ILogger
}

func NewUseCase(logger ILogger, repo IRepository) *useCaseImpl {
	return &useCaseImpl{
		repo:   repo,
		logger: logger,
	}
}

func (u *useCaseImpl) GetAllBike(ctx context.Context) ([]domain.Bike, error) {
	u.logger.Info("[BikeUseCase.GetAllBike] fetching all bikes")
	bike, err := u.repo.GetList(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []domain.Bike{}, nil
	}
	if err != nil {
		u.logger.Error("[BikeUseCase.GetAllBike] fetch all bikes failed", err)
		return []domain.Bike{}, apperrors.ErrInternalServerError
	}
	u.logger.Info("[BikeUseCase.GetAllBike] fetch all bikes success")
	return *bike, nil
}

func (u *useCaseImpl) Rent(ctx context.Context, payload domain.RentOrReturnRequestPayload) (*domain.Bike, error) {
	u.logger.Infof("[BikeUseCase.Rent] user %d is renting bike %d", payload.UserID, payload.ID)
	currentBike, err := u.repo.GetByID(ctx, payload.ID)
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
	err = u.repo.Update(ctx, updatedBike)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Rent] user %d rent bike %d failed", payload.UserID, payload.ID, err)
		return nil, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Rent] user %d rent bike %d success", payload.UserID, payload.ID)
	return updatedBike, nil
}

func (u *useCaseImpl) Return(ctx context.Context, payload domain.RentOrReturnRequestPayload) (*domain.Bike, error) {
	currentBike, err := u.repo.GetByID(ctx, payload.ID)
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
	err = u.repo.Update(ctx, updatedBike)
	if err != nil {
		u.logger.Errorf("[BikeUseCase.Return] user %d is return bike %d failed", currentBike.UserID, payload.ID, err)
		return nil, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[BikeUseCase.Return] user %d is return bike %d failed", currentBike.UserID, payload.ID)
	return updatedBike, nil
}
