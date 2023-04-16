package v1

import (
	"context"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// CustomerEndpoints - конечные точки для работы с клиентами
type CustomerEndpoints struct {
	GetCustomersList global.Endpoint
	GetCustomerByID  global.Endpoint
	CustomerStore    global.Endpoint
}

func (eps *CustomerEndpoints) makeCustomerEndpoints(s service.Service, middlewares ...global.Middleware) {
	eps.GetCustomersList = makeGetCustomersListEndpoint(s)
	eps.GetCustomerByID = makeGetCustomerByIDEndpoint(s)
	eps.CustomerStore = makeCustomerStoreEndpoint(s)

	// применяем промежуточное ПО
	for _, m := range middlewares {
		eps.GetCustomersList = m(eps.GetCustomersList)
		eps.GetCustomerByID = m(eps.GetCustomerByID)
		eps.CustomerStore = m(eps.CustomerStore)
	}
}

func makeGetCustomersListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestCustomerListFilters)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		domains, err := s.GetListCustomers(ctx, req.FioFilter, req.IsArchive, req.ClaimID)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.Customer, len(domains))
		for i := range domains {
			dtoRes[i] = toCustomerDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}

func makeGetCustomerByIDEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestByID)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		customer, err := s.GetCustomerByID(ctx, req.ID)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(toCustomerDTO(customer)), nil
	}
}

func makeCustomerStoreEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.Customer)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		id, err := s.StoreCustomer(ctx, fromCustomerDTO(req))
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(id), nil
	}
}
