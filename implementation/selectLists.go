package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) GetSelectListEmployees(ctx context.Context) ([]domain.SelectList, error) {
	employeesLst, err := bs.rep.SelectSelectListEmployees(ctx)
	if err != nil {
		log.Println("GetSelectListEmployees err:", err)
		return nil, global.InternalServerErr
	}

	return employeesLst, nil
}

func (bs *BasicService) GetSelectListPosition(ctx context.Context) ([]domain.SelectList, error) {
	positionsLst, err := bs.rep.SelectSelectListPositions(ctx)
	if err != nil {
		log.Println("GetSelectListPosition err:", err)
		return nil, global.InternalServerErr
	}

	return positionsLst, nil
}

func (bs *BasicService) GetSelectListLevelAccesses(ctx context.Context) ([]domain.LevelAccess, error) {
	levelAccessesLst, err := bs.rep.SelectSelectListLevelAccesses(ctx)
	if err != nil {
		log.Println("GetSelectListLevelAccesses err:", err)
		return nil, global.InternalServerErr
	}

	return levelAccessesLst, nil
}

func (bs *BasicService) GetSelectListFreedomType(ctx context.Context) ([]domain.SelectList, error) {
	freedomTypesLst, err := bs.rep.SelectSelectListFreedomTypes(ctx)
	if err != nil {
		log.Println("GetSelectListFreedomType err:", err)
		return nil, global.InternalServerErr
	}

	return freedomTypesLst, nil
}
