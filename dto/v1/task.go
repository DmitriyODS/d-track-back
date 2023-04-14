package v1

type RequestTaskListFilters struct {
	IsArchive    bool   `form:"is_archive"`
	NumberFilter string `form:"number_filter"`
}

type Task struct {
	ID                      uint64     `json:"id"`
	Number                  string     `json:"number"`
	DateCreated             int64      `json:"date_created"`
	DateCompleted           int64      `json:"date_completed"`
	DateEstimatedCompletion int64      `json:"date_estimated_completion"`
	Name                    string     `json:"name"`
	Description             string     `json:"description"`
	Status                  SelectList `json:"status"`
	Creator                 SelectList `json:"creator"`
	Executor                SelectList `json:"executor"`
}
