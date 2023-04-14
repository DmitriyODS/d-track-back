package store

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.com/ddda/d-track/d-track-back/domain"
	"gitlab.com/ddda/d-track/d-track-back/global"
)

const (
	selectTasksQuery = `
SELECT task.id,
       task.number,
       creator.id,
       creator.fio,
       task."name",
       task.description,
       task.date_created,
       task.date_completed,
       task.date_estimated_completion,
       ts.id,
       ts.name,
       executor.id,
       executor.fio
FROM user_data.tasks AS task
         INNER JOIN user_data.employees AS creator on creator.id = task.creator_id
         INNER JOIN user_data.task_states AS ts on ts.id = task.state_id
         INNER JOIN user_data.employees AS executor on executor.id = task.executor_id
`
	selectTaskByIDQuery = `
SELECT task.id,
       task.number,
       creator.id,
       creator.fio,
       task."name",
       task.description,
       task.date_created,
       task.date_completed,
       task.date_estimated_completion,
       ts.id,
       ts.name,
       executor.id,
       executor.fio
FROM user_data.tasks AS task
         INNER JOIN user_data.employees AS creator on creator.id = task.creator_id
         INNER JOIN user_data.task_states AS ts on ts.id = task.state_id
         INNER JOIN user_data.employees AS executor on executor.id = task.executor_id
WHERE task.id = $1
`
	createTaskQuery = `
INSERT INTO user_data.tasks(number,
                            creator_id,
                            "name",
                            description,
                            date_created,
                            date_estimated_completion,
                            state_id,
                            executor_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;
`
	updateTaskQuery = `
UPDATE user_data.tasks
SET "name"=$1,
    description=$2,
    date_estimated_completion=$3,
    state_id=$4,
    executor_id=$5
WHERE id = $6
RETURNING id;
`
)

func selectTaskPlaceholder(task *domain.Task) []interface{} {
	return []interface{}{
		&task.ID,
		&task.Number,
		&task.Creator.ID,
		&task.Creator.FIO,
		&task.Name,
		&task.Description,
		&task.DateCreated,
		&task.DateCompleted,
		&task.DateEstimatedCompletion,
		&task.Status.ID,
		&task.Status.Value,
		&task.Executor.ID,
		&task.Executor.FIO,
	}
}

func createTaskPlaceholder(task domain.Task) []interface{} {
	return []interface{}{
		task.Number,
		task.Creator.ID,
		task.Name,
		task.Description,
		task.DateCreated,
		task.DateEstimatedCompletion,
		task.Status.ID,
		task.Executor.ID,
	}
}
func updateTaskPlaceholder(task domain.Task) []interface{} {
	return []interface{}{
		task.Name,
		task.Description,
		task.DateEstimatedCompletion,
		task.Status.ID,
		task.Executor.ID,
		task.ID,
	}
}

func (s *Store) SelectTasks(ctx context.Context, numberFilter string, isArchive bool) ([]domain.Task, error) {
	sqlWithFilters := ""
	if isArchive {
		sqlWithFilters = fmt.Sprintf("%s WHERE ts.id=%d", selectTasksQuery, global.TaskStateClose)
	} else {
		sqlWithFilters = fmt.Sprintf("%s WHERE ts.id!=%d", selectTasksQuery, global.TaskStateClose)
	}

	if numberFilter != "" {
		sqlWithFilters = fmt.Sprintf("%s AND task.number ILIKE '%s'", sqlWithFilters, "%"+numberFilter+"%")
	}

	sqlWithFilters = fmt.Sprintf("%s ORDER by task.date_created DESC", sqlWithFilters)

	rows, err := s.Query(ctx, sqlWithFilters)
	if err == sql.ErrNoRows {
		return []domain.Task{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasksLst := make([]domain.Task, 0)

	var task domain.Task
	for rows.Next() {
		if err = rows.Scan(selectTaskPlaceholder(&task)...); err != nil {
			return nil, err
		}

		tasksLst = append(tasksLst, task)
	}

	return tasksLst, rows.Err()
}

func (s *Store) SelectTaskByID(ctx context.Context, id uint64) (domain.Task, error) {
	task := domain.NewTask(0)

	row, err := s.QueryRow(ctx, selectTaskByIDQuery, id)
	if err != nil {
		return task, err
	}

	err = row.Scan(selectTaskPlaceholder(&task)...)
	if err == sql.ErrNoRows {
		return task, nil
	}

	return task, err
}

func (s *Store) CreateTask(ctx context.Context, task domain.Task) (uint64, error) {
	row, err := s.QueryRow(ctx, createTaskQuery, createTaskPlaceholder(task)...)
	if err != nil {
		return 0, err
	}

	var resId uint64
	err = row.Scan(&resId)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	return resId, err
}

func (s *Store) UpdateTask(ctx context.Context, task domain.Task) (uint64, error) {
	var row *sql.Row
	var err error

	row, err = s.QueryRow(ctx, updateTaskQuery, updateTaskPlaceholder(task)...)
	if err != nil {
		return 0, err
	}

	var resId uint64
	err = row.Scan(&resId)
	if err == sql.ErrNoRows {
		return 0, nil
	}

	return resId, err
}
