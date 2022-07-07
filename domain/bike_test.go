package domain

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type BikeDomainTestSuite struct {
	suite.Suite
	bike *Bike
}

func (s *BikeDomainTestSuite) SetupTest() {
	mockTime := time.Time{}
	lat := decimal.NewFromFloat(50.119504)
	long := decimal.NewFromFloat(8.638137)
	bike := Bike{
		ID:        1,
		Lat:       &lat,
		Long:      &long,
		Status:    BikeStatusAvailable,
		UserID:    nil,
		CreatedAt: mockTime,
		UpdatedAt: mockTime,
		DeletedAt: gorm.DeletedAt{Valid: false},
	}
	s.bike = &bike
}

func (s *BikeDomainTestSuite) TestToDTO_Success() {
	actual := s.bike.ToDTO()
	expected := BikeDTO{
		ID:     s.bike.ID,
		Lat:    s.bike.Lat.String(),
		Long:   s.bike.Long.String(),
		Status: s.bike.Status,
		UserID: s.bike.UserID,
	}
	s.Equal(expected, actual)
}

func TestBikeDomainTestSuite(t *testing.T) {
	suite.Run(t, new(BikeDomainTestSuite))
}

func (s *BikeDomainTestSuite) TestTableName_Success() {
	tableName := s.bike.TableName()
	s.Equal("bike", tableName)
}

func (s *BikeDomainTestSuite) TestIsAvailable_Success() {
	isAvailable := s.bike.IsAvailable()
	isRented := s.bike.IsRented()
	s.True(isAvailable)
	s.False(isRented)
}

func (s *BikeDomainTestSuite) TestIsRented_Success() {
	s.bike.Status = BikeStatusRented
	userID := int64(1)
	s.bike.UserID = &userID
	isRent := s.bike.IsRented()
	isAvailable := s.bike.IsAvailable()
	s.True(isRent)
	s.False(isAvailable)
}
