package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type auth interface {
	AuthenticationByLogin(ctx context.Context, auth domain.Auth) (domain.Auth, error)
	CheckLevelAccessByEmployeeID(ctx context.Context, id uint64, levelAccess []byte) (bool, error)
}
