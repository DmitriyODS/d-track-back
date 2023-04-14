package domain

import "time"

type Task struct {
	ID                      uint64
	Number                  string
	Creator                 Employee
	Name                    string
	Description             string
	DateCreated             time.Time
	DateCompleted           time.Time
	DateEstimatedCompletion time.Time
	Status                  SelectList
	Executor                Employee
}

func NewTask(id uint64) Task {
	return Task{
		ID:                      id,
		Number:                  "",
		Creator:                 NewEmployee(0),
		Name:                    "",
		Description:             "",
		DateCreated:             time.Time{},
		DateCompleted:           time.Time{},
		DateEstimatedCompletion: time.Time{},
		Status:                  SelectList{},
		Executor:                NewEmployee(0),
	}
}

func (t Task) ValidateFields(isCreate bool) bool {
	return true
}

func (t Task) DateCreatedUnix() int64 {
	if t.DateCreated.IsZero() {
		return 0
	}

	return t.DateCreated.Unix()
}

func (t Task) DateCompletedUnix() int64 {
	if t.DateCompleted.IsZero() {
		return 0
	}

	return t.DateCompleted.Unix()
}

func (t Task) DateEstimatedCompletionUnix() int64 {
	if t.DateEstimatedCompletion.IsZero() {
		return 0
	}

	return t.DateEstimatedCompletion.Unix()
}