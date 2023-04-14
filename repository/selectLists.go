package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type repositorySelectLists interface {
	SelectSelectListEmployees(ctx context.Context) ([]domain.SelectList, error)
	SelectSelectListCustomers(ctx context.Context) ([]domain.SelectList, error)
	SelectSelectListPositions(ctx context.Context) ([]domain.SelectList, error)
	SelectSelectListLevelAccesses(ctx context.Context) ([]domain.LevelAccess, error)
	SelectSelectListFreedomTypes(ctx context.Context) ([]domain.SelectList, error)
	SelectSelectListServices(ctx context.Context) ([]domain.SelectList, error)
	SelectSelectListClaimStates(ctx context.Context) ([]domain.SelectList, error)
	SelectSelectListTaskStates(ctx context.Context) ([]domain.SelectList, error)
}
