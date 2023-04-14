package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type selectLists interface {
	GetSelectListEmployees(ctx context.Context) ([]domain.SelectList, error)
	GetSelectListCustomers(ctx context.Context) ([]domain.SelectList, error)
	GetSelectListPosition(ctx context.Context) ([]domain.SelectList, error)
	GetSelectListLevelAccesses(ctx context.Context) ([]domain.LevelAccess, error)
	GetSelectListFreedomType(ctx context.Context) ([]domain.SelectList, error)
	GetSelectListServices(ctx context.Context) ([]domain.SelectList, error)
	GetSelectListClaimStates(ctx context.Context) ([]domain.SelectList, error)
	GetSelectListTaskStates(ctx context.Context) ([]domain.SelectList, error)
}
