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

// ValidateFields проверяет корректность заполнения полей:
// - номер заявки не может быть пустым
// - дата создания не может быть пустой
// - ориентировочная дата выполнения не может быть пустой
// - заказчик не может быть пустым
// - тема не может быть пустой
// - тип услуги не может быть пустым
// - статус не может быть пустым
// - исполнитель не может быть пустым
func (c Claim) ValidateFields(_ bool) bool {
	if c.Number == "" {
		return false
	}

	if c.DateCreated.IsZero() {
		return false
	}

	if c.DateEstimatedCompletion.IsZero() {
		return false
	}

	if c.Customer.ID == 0 {
		return false
	}

	if c.Subject == "" {
		return false
	}

	if c.ServiceType.ID == 0 {
		return false
	}

	if c.Status.ID == 0 {
		return false
	}

	if c.Executor.ID == 0 {
		return false
	}

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
