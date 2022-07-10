package user

import (
	"context"

	"shared-bike/domain"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) GetListByIDs(ctx context.Context, IDs []int64) (*[]domain.User, error) {
	user := []domain.User{}
	err := r.db.Where("id IN (?)", IDs).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) Create(ctx context.Context, body *domain.User) error {
	err := r.db.Create(body).Error
	if err != nil {
		return err
	}
	return nil
}
