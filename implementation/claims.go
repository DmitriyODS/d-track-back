package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

func (bs *BasicService) GetListClaims(ctx context.Context, fioFilter string, isArchive bool) ([]domain.Claim, error) {
	return nil, nil
}

func (bs *BasicService) GetClaimByID(ctx context.Context, id uint64) (domain.Claim, error) {
	return domain.Claim{}, nil
}

func (bs *BasicService) StoreClaim(ctx context.Context, claim domain.Claim) (uint64, error) {
	return 0, nil
}
