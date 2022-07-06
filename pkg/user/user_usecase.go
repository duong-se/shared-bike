package user

import (
	"context"
	"errors"

	"github.com/duong-se/shared-bike/apperrors"
	"github.com/duong-se/shared-bike/domain"
	"gorm.io/gorm"
)

type useCaseImpl struct {
	repository IRepository
	logger     ILogger
}

func NewUseCase(logger ILogger, repository IRepository) *useCaseImpl {
	return &useCaseImpl{
		logger:     logger,
		repository: repository,
	}
}

func (u *useCaseImpl) Login(ctx context.Context, payload domain.LoginPayload) (domain.User, error) {
	u.logger.Info("[UserUseCase.Login] starting")
	user, err := u.repository.GetByUsername(ctx, payload.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Info("[UserUseCase.Login] user not found")
		return domain.User{}, apperrors.ErrUserNotFound
	}
	if err != nil {
		u.logger.Error("[UserUseCase.Login] fetch user by username failed", err)
		return domain.User{}, apperrors.ErrInternalServerError
	}
	if !user.ValidatePassword(payload.Password) {
		u.logger.Infof("[UserUseCase.Login] user %d login with password does not match", user.ID)
		return domain.User{}, apperrors.ErrUserNotFound
	}
	u.logger.Infof("[UserUseCase.Login] user %d login success", user.ID)
	return *user, nil
}

func (u *useCaseImpl) Register(ctx context.Context, payload domain.RegisterPayload) (domain.User, error) {
	u.logger.Info("[UserUseCase.Register] starting")
	newUser := domain.User{
		Username: payload.Username,
		Name:     payload.Name,
	}
	hashedPassword, err := newUser.HashPassword(payload.Password)
	if err != nil {
		u.logger.Error("[UserUseCase.Register] hash password failed", err)
		return domain.User{}, apperrors.ErrInternalServerError
	}
	newUser.Password = hashedPassword
	err = u.repository.Create(ctx, &newUser)
	if err != nil {
		u.logger.Error("[UserUseCase.Register] register failed", err)
		return domain.User{}, apperrors.ErrInternalServerError
	}
	u.logger.Infof("[UserUseCase.Register] user %d register success", newUser.ID)
	return newUser, nil
}
