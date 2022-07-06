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

type GetAllBikeResponse struct {
	ID               int64      `json:"id"`
	Lat              string     `json:"lat"`
	Long             string     `json:"long"`
	Status           BikeStatus `json:"status"`
	UserID           *int64     `json:"userId"`
	NameOfRenter     *string    `json:"nameOfRenter"`
	UsernameOfRenter *string    `json:"usernameOfRenter"`
}
