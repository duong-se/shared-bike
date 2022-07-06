package user

import (
	"context"

	"github.com/duong-se/shared-bike/domain"
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

func (r *repositoryImpl) Create(ctx context.Context, payload *domain.User) error {
	err := r.db.Create(payload).Error
	if err != nil {
		return err
	}
	return nil
}
