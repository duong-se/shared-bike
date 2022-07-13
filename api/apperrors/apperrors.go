package apperrors

import (
	"errors"
	"net/http"
)

var (
	// 500
	ErrInternalServerError = errors.New("e5000 internal server error")
	// 401
	ErrUnauthorizeError = errors.New("e4010 unauthorized")
	// 400
	ErrBikeRented         = errors.New("e4000 cannot rent because the bike is rented")
	ErrUserHasBikeAlready = errors.New("e4001 cannot rent because you have already rented a bike")
	ErrBikeAvailable      = errors.New("e4002 cannot return because the bike is available")
	ErrBikeNotYours       = errors.New("e4003 cannot return because the bike is not yours")
	ErrUserAlreadyExisted = errors.New("e4004 user already existed")
	ErrInvalidBody        = errors.New("e4005 invalid body")
	ErrInvalidBikeID      = errors.New("e4006 invalid bike id")
	// 404
	ErrBikeNotFound      = errors.New("e4040 bike not found")
	ErrUserLoginNotFound = errors.New("e4041 username or password is wrong")
	ErrUserNotExisted    = errors.New("e4042 user does not exist or inactive")
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
	case ErrInvalidBody:
		return http.StatusBadRequest
	case ErrInvalidBikeID:
		return http.StatusBadRequest
	case ErrUserAlreadyExisted:
		return http.StatusBadRequest
	case ErrUserLoginNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
