package v1

import (
	"context"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// EmployeeEndpoints - конечные точки для работы с сотрудниками
type EmployeeEndpoints struct {
	GetEmployeesList global.Endpoint
	GetEmployeeByID  global.Endpoint
	EmployeeStore    global.Endpoint
}

func (eps *EmployeeEndpoints) makeEmployeeEndpoints(s service.Service, middlewares ...global.Middleware) {
	eps.GetEmployeesList = makeGetEmployeesListEndpoint(s)
	eps.GetEmployeeByID = makeGetEmployeeByIDEndpoint(s)
	eps.EmployeeStore = makeEmployeeStoreEndpoint(s)

	// применяем промежуточное ПО
	for _, m := range middlewares {
		eps.GetEmployeesList = m(eps.GetEmployeesList)
		eps.GetEmployeeByID = m(eps.GetEmployeeByID)
		eps.EmployeeStore = m(eps.EmployeeStore)
	}
}

func makeGetEmployeesListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestList)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		domains, err := s.GetListEmployees(ctx, req.Filters, req.Sorts)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.Employee, len(domains))
		for i := range domains {
			dtoRes[i] = toEmployeeDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}

func makeGetEmployeeByIDEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestByID)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		employee, err := s.GetEmployeeByID(ctx, req.ID)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(toEmployeeDTO(employee)), nil
	}
}

func makeEmployeeStoreEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.Employee)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		id, err := s.StoreEmployee(ctx, fromEmployeeDTO(req))
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(id), nil
	}
}
