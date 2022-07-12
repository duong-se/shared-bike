package apperrors

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrUnauthorizeError    = errors.New("unauthorized")
	ErrBikeNotFound        = errors.New("bike not found")
	ErrBikeRented          = errors.New("cannot rent because bike is rented")
	ErrUserHasBikeAlready  = errors.New("cannot rent because you have already rented a bike")
	ErrBikeAvailable       = errors.New("cannot return because bike is available")
	ErrBikeNotYours        = errors.New("cannot return because bike is not yours")
	ErrUserLoginNotFound   = errors.New("username or password is wrong")
	ErrUserNotExisted      = errors.New("user not exists or inactive")
	ErrUserAlreadyExisted  = errors.New("user already existed")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrUnauthorizeError:
		return http.StatusUnauthorized
	case ErrBikeNotFound:
		return http.StatusNotFound
	case ErrBikeRented:
		return http.StatusBadRequest
	case ErrUserHasBikeAlready:
		return http.StatusBadRequest
	case ErrBikeAvailable:
		return http.StatusBadRequest
	case ErrUserNotExisted:
		return http.StatusBadRequest
	case ErrUserAlreadyExisted:
		return http.StatusBadRequest
	case ErrUserLoginNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
