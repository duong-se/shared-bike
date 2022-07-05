package bike

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

func (r *repositoryImpl) GetList(ctx context.Context) (*[]domain.Bike, error) {
	bikes := []domain.Bike{}
	err := r.db.Find(&bikes).Error
	if err != nil {
		return nil, err
	}
	return &bikes, nil
}

func (r *repositoryImpl) GetByID(ctx context.Context, id int64) (*domain.Bike, error) {
	bike := domain.Bike{}
	err := r.db.Where("id = ?", id).First(&bike).Error
	if err != nil {
		return nil, err
	}
	return &bike, nil
}

func (r *repositoryImpl) Update(ctx context.Context, payload *domain.Bike) error {
	err := r.db.Where("id = ?", payload.ID).Updates(payload).Error
	if err != nil {
		return err
	}
	return nil
}
