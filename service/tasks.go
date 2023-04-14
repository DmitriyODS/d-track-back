package service

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type tasks interface {
	GetListTasks(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, id uint64) (domain.Task, error)
	StoreTask(ctx context.Context, task domain.Task) (uint64, error)
}
