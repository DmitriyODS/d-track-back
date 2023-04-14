package v1

import (
	"gitlab.com/ddda/d-track/d-track-back/domain"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
	"time"
)

// Endpoints - объект конечных точек, который связывает все endpoints из пакета
// здесь встраиваем все endpoints - структуры
type Endpoints struct {
	AuthEndpoints
	EmployeeEndpoints
	SelectListEndpoints
	ClaimEndpoints
}

// MakeEndpoints - функция возвращает объект с конечными точками
// здесь создаём все endpoints, которые будут обслуживаться нашим сервисом
// middlewares - мидл вари, которые будут оборачивать доступ к конечным точкам
// s - сервис, который будет снабжать бизнес логикой конечные точки
func MakeEndpoints(s service.Service, middlewares ...global.Middleware) Endpoints {
	eps := Endpoints{}

	eps.makeAuthEndpoints(s, middlewares...)
	eps.makeEmployeeEndpoints(s, middlewares...)
	eps.makeSelectListEndpoints(s, middlewares...)
	eps.makeClaimEndpoints(s, middlewares...)

	return eps
}

func toSelectListsDTO(l domain.SelectList) dtoV1.SelectList {
	return dtoV1.SelectList{
		ID:    l.ID,
		Value: l.Value,
	}
}

func fromSelectListsDTO(l dtoV1.SelectList) domain.SelectList {
	return domain.SelectList{
		ID:    l.ID,
		Value: l.Value,
	}
}

func toLevelAccessDTO(l domain.LevelAccess) dtoV1.LevelAccess {
	return dtoV1.LevelAccess{
		ID:     l.ID,
		Name:   l.Name,
		Access: l.Access,
	}
}

func fromLevelAccessDTO(l dtoV1.LevelAccess) domain.LevelAccess {
	return domain.LevelAccess{
		ID:     l.ID,
		Name:   l.Name,
		Access: l.Access,
	}
}

func toEmployeeDTO(e domain.Employee) dtoV1.Employee {
	return dtoV1.Employee{
		ID:                 e.ID,
		FIO:                e.FIO,
		Login:              e.Login,
		PhoneNumber:        e.PhoneNumber,
		EmailAddress:       e.EmailAddress,
		AddressOfResidence: e.AddressOfResidence,
		Position:           toSelectListsDTO(e.Position),
		LevelAccess:        toLevelAccessDTO(e.LevelAccess),
		FreedomType:        toSelectListsDTO(e.FreedomType),
		DateAppointments:   e.DateAppointmentsUnix(),
		DateOfDismissal:    e.DateOfDismissalUnix(),
	}
}

func fromEmployeeDTO(e dtoV1.Employee) domain.Employee {
	return domain.Employee{
		ID:                 e.ID,
		FIO:                e.FIO,
		Login:              e.Login,
		Password:           e.Password,
		PhoneNumber:        e.PhoneNumber,
		EmailAddress:       e.EmailAddress,
		AddressOfResidence: e.AddressOfResidence,
		Position:           fromSelectListsDTO(e.Position),
		LevelAccess:        fromLevelAccessDTO(e.LevelAccess),
		FreedomType:        fromSelectListsDTO(e.FreedomType),
		DateAppointments:   time.Unix(e.DateAppointments, 0),
		DateOfDismissal:    time.Unix(e.DateOfDismissal, 0),
	}
}

func toAuthDTO(l domain.Auth) dtoV1.Auth {
	return dtoV1.Auth{
		UserID: l.UserID,
		Login:  l.Login,
		JWT:    l.JWT,
	}
}

func fromAuthDTO(l dtoV1.Auth) domain.Auth {
	return domain.Auth{
		UserID:   l.UserID,
		Login:    l.Login,
		Password: l.Password,
		JWT:      l.JWT,
	}
}

func toCustomerDTO(c domain.Customer) dtoV1.Customer {
	return dtoV1.Customer{
		ID:          c.ID,
		FIO:         c.FIO,
		Phone:       c.Phone,
		Email:       c.Email,
		Address:     c.Address,
		DateCreated: c.DateCreatedUnix(),
	}
}

func fromCustomerDTO(c dtoV1.Customer) domain.Customer {
	return domain.Customer{
		ID:          c.ID,
		FIO:         c.FIO,
		Phone:       c.Phone,
		Email:       c.Email,
		Address:     c.Address,
		DateCreated: time.Unix(c.DateCreated, 0),
	}
}

func toClaimDTO(c domain.Claim) dtoV1.Claim {
	return dtoV1.Claim{
		ID:                      c.ID,
		Number:                  c.Number,
		DateCreated:             c.DateCreatedUnix(),
		DateCompleted:           c.DateCompletedUnix(),
		DateEstimatedCompletion: c.DateEstimatedCompletionUnix(),
		Customer: dtoV1.SelectList{
			ID:    c.Customer.ID,
			Value: c.Customer.FIO,
		},
		Subject:     c.Subject,
		ServiceType: toSelectListsDTO(c.ServiceType),
		Description: c.Description,
		Status:      toSelectListsDTO(c.Status),
		Executor: dtoV1.SelectList{
			ID:    c.Executor.ID,
			Value: c.Executor.FIO,
		},
	}
}

func fromClaimDTO(c dtoV1.Claim) domain.Claim {
	return domain.Claim{
		ID:                      c.ID,
		Number:                  c.Number,
		DateCreated:             time.Unix(c.DateCreated, 0),
		DateCompleted:           time.Unix(c.DateCompleted, 0),
		DateEstimatedCompletion: time.Unix(c.DateEstimatedCompletion, 0),
		Customer:                domain.NewCustomer(c.Customer.ID),
		Subject:                 c.Subject,
		ServiceType:             fromSelectListsDTO(c.ServiceType),
		Description:             c.Description,
		Status:                  fromSelectListsDTO(c.Status),
		Executor:                domain.NewEmployee(c.Executor.ID),
	}
}

func toTaskDTO(t domain.Task) dtoV1.Task {
	return dtoV1.Task{
		ID:                      t.ID,
		Number:                  t.Number,
		DateCreated:             t.DateCreatedUnix(),
		DateCompleted:           t.DateCompletedUnix(),
		DateEstimatedCompletion: t.DateEstimatedCompletionUnix(),
		Name:                    t.Name,
		Description:             t.Description,
		Status:                  toSelectListsDTO(t.Status),
		Creator: dtoV1.SelectList{
			ID:    t.Creator.ID,
			Value: t.Creator.FIO,
		},
		Executor: dtoV1.SelectList{
			ID:    t.Executor.ID,
			Value: t.Executor.FIO,
		},
	}
}

func fromTaskDTO(t dtoV1.Task) domain.Task {
	return domain.Task{
		ID:                      t.ID,
		Number:                  t.Number,
		Creator:                 domain.NewEmployee(t.Creator.ID),
		Name:                    t.Name,
		Description:             t.Description,
		DateCreated:             time.Unix(t.DateCreated, 0),
		DateCompleted:           time.Unix(t.DateCompleted, 0),
		DateEstimatedCompletion: time.Unix(t.DateEstimatedCompletion, 0),
		Status:                  fromSelectListsDTO(t.Status),
		Executor:                domain.NewEmployee(t.Executor.ID),
	}
}
