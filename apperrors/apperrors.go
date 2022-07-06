package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrBikeNotFound        = errors.New("your request bike not found")
	ErrBikeRented          = errors.New("cannot rent because bike is rented")
	ErrBikeAvailable       = errors.New("cannot return because bike is available")
	ErrBikeNotYours        = errors.New("cannot return because bike is not yours")
	ErrUserNotFound        = errors.New("username or password is wrong")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrBikeNotFound:
		return http.StatusNotFound
	case ErrBikeRented:
		return http.StatusBadRequest
	case ErrBikeAvailable:
		return http.StatusBadRequest
	case ErrUserNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}