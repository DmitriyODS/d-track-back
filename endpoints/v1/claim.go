package v1

import (
	"context"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// ClaimEndpoints - конечные точки для работы с заявками
type ClaimEndpoints struct {
	GetClaimsList global.Endpoint
	GetClaimByID  global.Endpoint
	ClaimStore    global.Endpoint
}

func (eps *ClaimEndpoints) makeClaimEndpoints(s service.Service, middlewares ...global.Middleware) {
	eps.GetClaimsList = makeGetClaimsListEndpoint(s)
	eps.GetClaimByID = makeGetClaimByIDEndpoint(s)
	eps.ClaimStore = makeClaimStoreEndpoint(s)

	// применяем промежуточное ПО
	for _, m := range middlewares {
		eps.GetClaimsList = m(eps.GetClaimsList)
		eps.GetClaimByID = m(eps.GetClaimByID)
		eps.ClaimStore = m(eps.ClaimStore)
	}
}

func makeGetClaimsListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestEmployeeListFilters)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		domains, err := s.GetListEmployees(ctx, req.FioFilter, req.IsArchive)
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

func makeGetClaimByIDEndpoint(s service.Service) global.Endpoint {
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

func makeClaimStoreEndpoint(s service.Service) global.Endpoint {
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
