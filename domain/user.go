package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterBody struct {
	Username string `json:"username" example:"myusername"`
	Password string `json:"password" example:"mypassword"`
	Name     string `json:"name" example:"myname"`
}

type LoginBody struct {
	Username string `json:"username" example:"myusername"`
	Password string `json:"password" example:"mypassword"`
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

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
	}
}

func (u *User) ValidatePassword(plainPassword string) bool {
	password := []byte(plainPassword)
	hashedPassword := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}

func (u *User) HashPassword(plainPassword string, cost int) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), cost)
	if err != nil {
		return "", fmt.Errorf("bycrpt password: %w", err)
	}
	return string(hashedPassword), nil
}

func (User) TableName() string {
	return "user"
}

type UserDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
