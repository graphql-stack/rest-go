package common

import (
	"errors"
	"github.com/zcong1993/libgo/gin/ginerr"
	"net/http"
)

func createErr(statusCode int, code string) ginerr.ApiError {
	return ginerr.NewDefaultError(statusCode, code, code)
}

var (
	DUPLICATE_USER               = ginerr.NewDefaultError(http.StatusBadRequest, "DUPLICATE_USER", "DUPLICATE_USER")
	INTERVAL_ERROR               = ginerr.NewDefaultError(http.StatusInternalServerError, "", "")
	INVALID_PARAMS               = ginerr.NewDefaultError(http.StatusBadRequest, "INVALID_PARAMS", "INVALID_PARAMS")
	INVALID_USERNAME_OR_PASSWORD = createErr(http.StatusUnauthorized, "INVALID_USERNAME_OR_PASSWORD")
)

var TOKEN_ERROR = map[string]string{"code": "TOKEN_ERR_OR_EXPIRED", "message": "TOKEN_ERR_OR_EXPIRED"}

var (
	ErrExpired = errors.New("TOKEN_EXPIRED")
)

type ErrResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
