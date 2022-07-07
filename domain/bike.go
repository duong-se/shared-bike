package domain

import (
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
	Lat       *decimal.Decimal `json:"lat"`
	Long      *decimal.Decimal `json:"long"`
	Status    BikeStatus       `json:"status"`
	UserID    *int64           `json:"userId"`
	CreatedAt time.Time        `json:"-"`
	UpdatedAt time.Time        `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-"`
}

func (b *Bike) ToDTO() BikeDTO {
	return BikeDTO{
		ID:     b.ID,
		Lat:    b.Lat.String(),
		Long:   b.Long.String(),
		Status: b.Status,
		UserID: b.UserID,
	}
}

func (b *Bike) IsRented() bool {
	return b.Status == BikeStatusRented && b.UserID != nil
}

func (b *Bike) IsAvailable() bool {
	return b.Status == BikeStatusAvailable && b.UserID == nil
}

func (Bike) TableName() string {
	return "bike"
}

type RentOrReturnRequestPayload struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"userId"`
}

type BikeDTO struct {
	ID               int64      `json:"id" example:"1"`
	Lat              string     `json:"lat" example:"50.119504"`
	Long             string     `json:"long" example:"8.638137"`
	Status           BikeStatus `json:"status" example:"rented"`
	UserID           *int64     `json:"userId" example:"1"`
	NameOfRenter     *string    `json:"nameOfRenter" example:"Bob"`
	UsernameOfRenter *string    `json:"usernameOfRenter" example:"bob"`
}
