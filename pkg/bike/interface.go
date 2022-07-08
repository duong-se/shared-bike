package bike

import (
	"context"

	"shared-bike/domain"
)

type IRepository interface {
	GetList(ctx context.Context) (*[]domain.Bike, error)
	GetByID(ctx context.Context, id int64) (*domain.Bike, error)
	UpdateStatusAndUserID(ctx context.Context, body *domain.Bike) error
	CountByUserID(ctx context.Context, id int64) (int64, error)
}

type IUserRepository interface {
	GetListByIDs(ctx context.Context, IDs []int64) (*[]domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
}

type ILogger interface {
	Info(i ...interface{})
	Warn(i ...interface{})
	Error(i ...interface{})
}

type IUseCase interface {
	GetAllBike(ctx context.Context) ([]domain.BikeDTO, error)
	Rent(ctx context.Context, body domain.RentOrReturnRequestPayload) (domain.BikeDTO, error)
	Return(ctx context.Context, body domain.RentOrReturnRequestPayload) (domain.BikeDTO, error)
}

//go:generate mockery --name IRepository --output mocks --case underscore
//go:generate mockery --name ILogger --output mocks --case underscore
//go:generate mockery --name IUseCase --output mocks --case underscore
//go:generate mockery --name IUserRepository --output mocks --case underscore
