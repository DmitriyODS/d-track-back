package v1

type RequestClaimListFilters struct {
	IsArchive bool   `form:"is_archive"`
	FioFilter string `form:"fio_filter"`
}

type Claim struct {
	ID                      uint64     `json:"id"`
	Number                  string     `json:"number"`
	DateCreated             int64      `json:"date_created"`
	DateCompleted           int64      `json:"date_completed"`
	DateEstimatedCompletion int64      `json:"date_estimated_completion"`
	Customer                SelectList `json:"customer"`
	Subject                 string     `json:"subject"`
	ServiceType             SelectList `json:"service_type"`
	Description             string     `json:"description"`
	Status                  SelectList `json:"status"`
	Executor                SelectList `json:"executor"`
}
