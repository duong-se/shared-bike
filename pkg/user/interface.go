package user

import (
	"context"

	"github.com/duong-se/shared-bike/domain"
)

type IRepository interface {
	GetByUsername(ctx context.Context, username string) (*domain.User, error)
	Create(ctx context.Context, payload *domain.User) error
}

type ILogger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Info(i ...interface{})
	Error(i ...interface{})
}
type IUseCase interface {
	Login(ctx context.Context, payload domain.LoginPayload) (domain.User, error)
	Register(ctx context.Context, payload domain.RegisterPayload) (domain.User, error)
}

//go:generate mockery --name IRepository --output mocks --case underscore
//go:generate mockery --name ILogger --output mocks --case underscore
//go:generate mockery --name IUseCase --output mocks --case underscore
