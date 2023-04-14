package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type claims interface {
	GetListClaims(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Claim, error)
	GetClaimByID(ctx context.Context, id uint64) (domain.Claim, error)
	StoreClaim(ctx context.Context, claim domain.Claim) (uint64, error)
}
