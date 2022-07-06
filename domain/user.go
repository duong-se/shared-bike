package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID        int64          `json:"id"`
	Username  string         `json:"username"`
	Password  string         `json:"-"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (u *User) ValidatePassword(plainPassword string) bool {
	password := []byte(plainPassword)
	hashedPassword := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}

func (u *User) HashPassword(plainPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bycrpt password: %w", err)
	}
	return string(hashedPassword), nil
}

func (User) TableName() string {
	return "user"
}
