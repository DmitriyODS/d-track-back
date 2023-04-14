package repository

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

type repositoryTasks interface {
	SelectTasks(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Task, error)
	SelectTaskByID(ctx context.Context, id uint64) (domain.Task, error)
	CreateTask(ctx context.Context, task domain.Task) (uint64, error)
	UpdateTask(ctx context.Context, task domain.Task) (uint64, error)
}
