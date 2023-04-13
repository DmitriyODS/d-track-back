package domain

import "time"

type Customer struct {
	ID          uint64
	FIO         string
	Phone       string
	Email       string
	Address     string
	DateCreated time.Time
}

func NewCustomer(id uint64) Customer {
	return Customer{
		ID:          id,
		FIO:         "",
		Phone:       "",
		Email:       "",
		Address:     "",
		DateCreated: time.Time{},
	}
}

func (c Customer) ValidateFields(isCreate bool) bool {
	return true
}

func (c Customer) DateCreatedUnix() int64 {
	if c.DateCreated.IsZero() {
		return 0
	}

	return c.DateCreated.Unix()
}
