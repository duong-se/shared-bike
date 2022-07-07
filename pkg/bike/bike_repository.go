package bike

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

func (r *repositoryImpl) CountByUserID(ctx context.Context, id int64) (int64, error) {
	var total int64
	err := r.db.Model(domain.Bike{}).Where("user_id = ?", id).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repositoryImpl) Update(ctx context.Context, body *domain.Bike) error {
	err := r.db.Where("id = ?", body.ID).Updates(body).Error
	if err != nil {
		return err
	}
	return nil
}
