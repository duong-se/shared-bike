package bike

import (
	"context"

	"github.com/duong-se/shared-bike/domain"
)

type IRepository interface {
	GetList(ctx context.Context) (*[]domain.Bike, error)
	GetByID(ctx context.Context, id int64) (*domain.Bike, error)
	Update(ctx context.Context, payload *domain.Bike) error
}

type ILogger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Info(i ...interface{})
	Error(i ...interface{})
}

type IUseCase interface {
	GetAllBike(ctx context.Context) ([]domain.Bike, error)
	Rent(ctx context.Context, payload domain.RentOrReturnRequestPayload) (*domain.Bike, error)
	Return(ctx context.Context, payload domain.RentOrReturnRequestPayload) (*domain.Bike, error)
}

//go:generate mockery --name IRepository --output mocks --case underscore
//go:generate mockery --name ILogger --output mocks --case underscore
//go:generate mockery --name IUseCase --output mocks --case underscore
