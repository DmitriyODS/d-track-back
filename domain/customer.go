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

// ValidateFields проверяет корректность заполнения полей
// - все поля обязательны
func (c Customer) ValidateFields(_ bool) bool {
	if c.FIO == "" || c.Phone == "" || c.Email == "" || c.Address == "" {
		return false
	}

	if c.DateCreated.IsZero() {
		return false
	}

	return true
}

func (c Customer) DateCreatedUnix() int64 {
	if c.DateCreated.IsZero() {
		return 0
	}

	return c.DateCreated.Unix()
}
