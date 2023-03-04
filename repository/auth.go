package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type auth interface {
	SelectUserByLoginPass(ctx context.Context, auth domain.Auth) (domain.Employee, error)
	CheckLevelAccessByEmployeeID(ctx context.Context, id uint64, levelAccess []byte) (bool, error)
}
