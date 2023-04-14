package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
)

func (bs *BasicService) GetListTasks(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Task, error) {
	return nil, nil
}

func (bs *BasicService) GetTaskByID(ctx context.Context, id uint64) (domain.Task, error) {
	return domain.Task{}, nil
}

func (bs *BasicService) StoreTask(ctx context.Context, task domain.Task) (uint64, error) {
	return 0, nil
}
