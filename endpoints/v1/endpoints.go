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
