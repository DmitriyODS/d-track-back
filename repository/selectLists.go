package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type repositorySelectLists interface {
	SelectSelectListEmployees(ctx context.Context, filters, sorts map[string]string) ([]domain.SelectList, error)
	SelectSelectListPositions(ctx context.Context, filters, sorts map[string]string) ([]domain.SelectList, error)
	SelectSelectListLevelAccesses(ctx context.Context, filters, sorts map[string]string) ([]domain.LevelAccess, error)
	SelectSelectListFreedomTypes(ctx context.Context, filters, sorts map[string]string) ([]domain.SelectList, error)
}
