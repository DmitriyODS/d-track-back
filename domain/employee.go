package domain

import (
	"gitlab.com/ddda/d-track/d-track-back/global"
	"time"
)

type Employee struct {
	ID                 uint64
	FIO                string
	Login              string
	Password           string
	PhoneNumber        string
	EmailAddress       string
	AddressOfResidence string
	Position           SelectList
	LevelAccess        LevelAccess
	FreedomType        SelectList
	DateAppointments   time.Time
	DateOfDismissal    time.Time
}

func NewEmployee(id uint64) Employee {
	return Employee{
		ID:                 id,
		FIO:                "",
		Login:              "",
		Password:           "",
		PhoneNumber:        "",
		EmailAddress:       "",
		AddressOfResidence: "",
		Position:           SelectList{},
		LevelAccess:        LevelAccess{},
		FreedomType:        SelectList{},
		DateAppointments:   time.Time{},
		DateOfDismissal:    time.Time{},
	}
}

// IsDismissal проверяет, уволен ли сотрудник
func (e Employee) IsDismissal() bool {
	return e.FreedomType.ID == global.EmployeeFreedomTypeFired
}

func (e Employee) ValidateFields(isCreate bool) bool {
	if e.FIO == "" {
		return false
	}

	if e.Login == "" {
		return false
	}

	if e.Password == "" && isCreate {
		return false
	}

	if e.LevelAccess.ID == 0 {
		return false
	}

	if e.DateAppointments.IsZero() {
		return false
	}

	return true
}

func (e Employee) DateAppointmentsUnix() int64 {
	if e.DateAppointments.IsZero() {
		return 0
	}

	return e.DateAppointments.Unix()
}

func (e Employee) DateOfDismissalUnix() int64 {
	if e.DateOfDismissal.IsZero() {
		return 0
	}

	return e.DateOfDismissal.Unix()
}
