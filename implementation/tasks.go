package implementation

import (
	"context"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
)

func (bs *BasicService) GetListTasks(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Task, error) {
	tasks, err := bs.rep.SelectTasks(ctx, numberFilter, isArchive)
	if err != nil {
		log.Println("GetListTasks err:", err)
		return nil, global.InternalServerErr
	}

	return tasks, nil
}

func (bs *BasicService) GetTaskByID(ctx context.Context, id uint64) (domain.Task, error) {
	if id == 0 {
		return domain.NewTask(0), global.BadRequestErr
	}

	task, err := bs.rep.SelectTaskByID(ctx, id)
	if err != nil {
		log.Println("GetTaskByID err:", err)
		return task, global.InternalServerErr
	}

	if task.ID == 0 {
		return task, global.DataNotFoundErr
	}

	return task, nil
}

func (bs *BasicService) StoreTask(ctx context.Context, task domain.Task) (uint64, error) {
	// если id не задан - запрос на создание
	if task.ID == 0 {
		// не прошли валидацию, отправляемся обратно
		if !task.ValidateFields(true) {
			return 0, global.IncorrectValidFormErr
		}

		id, err := bs.rep.CreateTask(ctx, task)
		if err != nil || id == 0 {
			log.Println("StoreTask create err:", err)
			return 0, global.InternalServerErr
		}

		return id, nil
	}

	// проверка корректнности заполнения полей
	if !task.ValidateFields(false) {
		return 0, global.IncorrectValidFormErr
	}

	id, err := bs.rep.UpdateTask(ctx, task)
	if err != nil || id == 0 {
		log.Println("StoreTask update err:", err)
		return 0, global.InternalServerErr
	}

	return id, nil
}
