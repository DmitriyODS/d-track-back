package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type repositoryClaims interface {
	SelectClaims(ctx context.Context, numberFilter string, isArchive bool, customerID uint64) ([]domain.Claim, error)
	SelectClaimByID(ctx context.Context, id uint64) (domain.Claim, error)
	CreateClaim(ctx context.Context, claim domain.Claim) (uint64, error)
	UpdateClaim(ctx context.Context, claim domain.Claim) (uint64, error)
}
