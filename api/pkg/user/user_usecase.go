package user

import (
	"context"
	"errors"
	"fmt"

	"shared-bike/apperrors"
	"shared-bike/domain"

	"golang.org/x/crypto/bcrypt"
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

func (u *useCaseImpl) Login(ctx context.Context, body domain.LoginBody) (domain.UserDTO, error) {
	u.logger.Info("[UserUseCase.Login] starting")
	user, err := u.repository.GetByUsername(ctx, body.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Info("[UserUseCase.Login] user not found")
		return domain.UserDTO{}, apperrors.ErrUserLoginNotFound
	}
	if err != nil {
		u.logger.Error("[UserUseCase.Login] fetch user by username failed", err)
		return domain.UserDTO{}, apperrors.ErrInternalServerError
	}
	if !user.ValidatePassword(body.Password) {
		u.logger.Info(fmt.Sprintf("[UserUseCase.Login] user %d login with password does not match", user.ID))
		return domain.UserDTO{}, apperrors.ErrUserLoginNotFound
	}
	u.logger.Info(fmt.Sprintf("[UserUseCase.Login] user %d login success", user.ID))
	return user.ToDTO(), nil
}

func (u *useCaseImpl) Register(ctx context.Context, body domain.RegisterBody) (domain.UserDTO, error) {
	existedUser, err := u.repository.GetByUsername(ctx, body.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("[UserUseCase.Register] fetch user by username failed", err)
		return domain.UserDTO{}, apperrors.ErrInternalServerError
	}
	if existedUser != nil {
		u.logger.Info("[UserUseCase.Register] user already existed")
		return domain.UserDTO{}, apperrors.ErrUserAlreadyExisted
	}
	u.logger.Info("[UserUseCase.Register] starting")
	newUser := domain.User{
		Username: body.Username,
		Name:     body.Name,
	}
	hashedPassword, _ := newUser.HashPassword(body.Password, bcrypt.DefaultCost)
	newUser.Password = hashedPassword
	err = u.repository.Create(ctx, &newUser)
	if err != nil {
		u.logger.Error("[UserUseCase.Register] register failed", err)
		return domain.UserDTO{}, apperrors.ErrInternalServerError
	}
	u.logger.Info(fmt.Sprintf("[UserUseCase.Register] user %d register success", newUser.ID))
	return newUser.ToDTO(), nil
}
