package global

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID         uint64 `json:"user_id,omitempty"`
	UserPositionID uint64 `json:"user_position_id,omitempty"`
	UserLogin      string `json:"user_login,omitempty"`
	LevelAccess    byte   `json:"level_access,omitempty"`
	jwt.RegisteredClaims
}

// ResponseData - структура ответа на любой запрос
// OK - статус запроса: выполнен/не выполнен
// CodeErr - код ошибки, если запрос не выполнен
// Description - описание ошибки, если запрос не выполнен
// Data - результат запроса, не пустая если запрос выполне
type ResponseData struct {
	OK          bool        `json:"ok"`
	CodeErr     int         `json:"code_err,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func NewResponseData(data interface{}) ResponseData {
	return ResponseData{
		OK:   true,
		Data: data,
	}
}

func NewErrResponseData(err error) ResponseData {
	return ResponseData{
		OK:          false,
		CodeErr:     GetStatusCodeByErr(err),
		Description: err.Error(),
	}
}

type Endpoint func(ctx context.Context, request interface{}) (response ResponseData, err error)

type Middleware func(Endpoint) Endpoint
