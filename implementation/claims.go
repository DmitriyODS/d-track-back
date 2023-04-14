package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) GetListClaims(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Claim, error) {
	claims, err := bs.rep.SelectClaims(ctx, numberFilter, isArchive)
	if err != nil {
		log.Println("GetListClaims err:", err)
		return nil, global.InternalServerErr
	}

	return claims, nil
}

func (bs *BasicService) GetClaimByID(ctx context.Context, id uint64) (domain.Claim, error) {
	if id == 0 {
		return domain.NewClaim(0), global.BadRequestErr
	}

	claim, err := bs.rep.SelectClaimByID(ctx, id)
	if err != nil {
		log.Println("GetClaimByID err:", err)
		return claim, global.InternalServerErr
	}

	if claim.ID == 0 {
		return claim, global.DataNotFoundErr
	}

	return claim, nil
}

func (bs *BasicService) StoreClaim(ctx context.Context, claim domain.Claim) (uint64, error) {
	// если id не задан - запрос на создание
	if claim.ID == 0 {
		// не прошли валидацию, отправляемся обратно
		if !claim.ValidateFields(true) {
			return 0, global.IncorrectValidFormErr
		}

		id, err := bs.rep.CreateClaim(ctx, claim)
		if err != nil || id == 0 {
			log.Println("StoreClaim create err:", err)
			return 0, global.InternalServerErr
		}

		return id, nil
	}

	// проверка корректнности заполнения полей
	if !claim.ValidateFields(false) {
		return 0, global.IncorrectValidFormErr
	}

	id, err := bs.rep.UpdateClaim(ctx, claim)
	if err != nil || id == 0 {
		log.Println("StoreClaim update err:", err)
		return 0, global.InternalServerErr
	}

	return id, nil
}
