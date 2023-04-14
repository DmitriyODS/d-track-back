package domain

import "time"

type Claim struct {
	ID                      uint64
	Number                  string
	DateCreated             time.Time
	DateCompleted           time.Time
	DateEstimatedCompletion time.Time
	Customer                Customer
	Subject                 string
	ServiceType             SelectList
	Description             string
	Status                  SelectList
	Executor                Employee
}

func NewClaim(id uint64) Claim {
	return Claim{
		ID:                      id,
		Number:                  "",
		DateCreated:             time.Now(),
		DateCompleted:           time.Time{},
		DateEstimatedCompletion: time.Time{},
		Customer:                NewCustomer(0),
		Subject:                 "",
		ServiceType:             SelectList{},
		Description:             "",
		Status:                  SelectList{},
		Executor:                NewEmployee(0),
	}
}

func (c Claim) ValidateFields(isCreate bool) bool {
	return true
}

func (c Claim) DateCreatedUnix() int64 {
	if c.DateCreated.IsZero() {
		return 0
	}

	return c.DateCreated.Unix()
}

func (c Claim) DateCompletedUnix() int64 {
	if c.DateCompleted.IsZero() {
		return 0
	}

	return c.DateCompleted.Unix()
}

func (c Claim) DateEstimatedCompletionUnix() int64 {
	if c.DateEstimatedCompletion.IsZero() {
		return 0
	}

	return c.DateEstimatedCompletion.Unix()
}
