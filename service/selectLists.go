package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type selectLists interface {
	GetSelectListEmployees(ctx context.Context, filters, sorts map[string]string) ([]domain.SelectList, error)
	GetSelectListPosition(ctx context.Context, filters, sorts map[string]string) ([]domain.SelectList, error)
	GetSelectListLevelAccesses(ctx context.Context, filters, sorts map[string]string) ([]domain.LevelAccess, error)
	GetSelectListFreedomType(ctx context.Context, filters, sorts map[string]string) ([]domain.SelectList, error)
}
