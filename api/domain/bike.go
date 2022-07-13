package domain

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BikeStatus string

var (
	BikeStatusRented    BikeStatus = "rented"
	BikeStatusAvailable BikeStatus = "available"
)

type Bike struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	Lat       *decimal.Decimal `json:"lat"`
	Long      *decimal.Decimal `json:"long"`
	Status    BikeStatus       `json:"status"`
	UserID    sql.NullInt64    `json:"userId"`
	CreatedAt time.Time        `json:"-"`
	UpdatedAt time.Time        `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-"`
}

func (b *Bike) ToDTO() BikeDTO {
	bikeDTO := BikeDTO{
		ID:     b.ID,
		Name:   b.Name,
		Lat:    b.Lat.String(),
		Long:   b.Long.String(),
		Status: b.Status,
	}
	if b.UserID.Valid {
		bikeDTO.UserID = b.UserID.Int64
	}
	return bikeDTO
}

func (b *Bike) IsRented() bool {
	return b.Status == BikeStatusRented && b.UserID.Valid
}

func (b *Bike) IsAvailable() bool {
	return b.Status == BikeStatusAvailable && !b.UserID.Valid
}

func (Bike) TableName() string {
	return "bike"
}

type RentOrReturnRequestPayload struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"userId"`
}

type BikeDTO struct {
	ID           int64      `json:"id" example:"1"`
	Name         string     `json:"name" example:"henry"`
	Lat          string     `json:"lat" example:"50.119504"`
	Long         string     `json:"long" example:"8.638137"`
	Status       BikeStatus `json:"status" example:"rented"`
	UserID       int64      `json:"userId" example:"1"`
	NameOfRenter string     `json:"nameOfRenter" example:"Bob"`
}
