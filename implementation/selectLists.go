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

func (bs *BasicService) GetSelectListCustomers(ctx context.Context) ([]domain.SelectList, error) {
	employeesLst, err := bs.rep.SelectSelectListCustomers(ctx)
	if err != nil {
		log.Println("GetSelectListCustomers err:", err)
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

func (bs *BasicService) GetSelectListServices(ctx context.Context) ([]domain.SelectList, error) {
	servicesLst, err := bs.rep.SelectSelectListServices(ctx)
	if err != nil {
		log.Println("GetSelectListServices err:", err)
		return nil, global.InternalServerErr
	}

	return servicesLst, nil
}

func (bs *BasicService) GetSelectListClaimStates(ctx context.Context) ([]domain.SelectList, error) {
	servicesLst, err := bs.rep.SelectSelectListClaimStates(ctx)
	if err != nil {
		log.Println("GetSelectListClaimStates err:", err)
		return nil, global.InternalServerErr
	}

	return servicesLst, nil
}

func (bs *BasicService) GetSelectListTaskStates(ctx context.Context) ([]domain.SelectList, error) {
	servicesLst, err := bs.rep.SelectSelectListTaskStates(ctx)
	if err != nil {
		log.Println("GetSelectListTaskStates err:", err)
		return nil, global.InternalServerErr
	}

	return servicesLst, nil
}
