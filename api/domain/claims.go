package domain

import "github.com/golang-jwt/jwt"

var (
	UserIDKey   = "id"
	UsernameKey = "username"
	NameKey     = "name"
)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.StandardClaims
}
