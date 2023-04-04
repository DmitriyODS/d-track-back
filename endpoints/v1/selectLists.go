package v1

import (
	"context"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// SelectListEndpoints - конечные точки для получения списков выбора
type SelectListEndpoints struct {
	GetEmployeesSelectList     global.Endpoint
	GetPositionsSelectList     global.Endpoint
	GetLevelAccessesSelectList global.Endpoint
	GetFreedomTypesSelectList  global.Endpoint
}

func (eps *SelectListEndpoints) makeSelectListEndpoints(s service.Service, middlewares ...global.Middleware) {
	eps.GetEmployeesSelectList = makeEmployeesSelectListEndpoint(s)
	eps.GetPositionsSelectList = makePositionsSelectListEndpoint(s)
	eps.GetLevelAccessesSelectList = makeLevelAccessesSelectListEndpoint(s)
	eps.GetFreedomTypesSelectList = makeFreedomTypesSelectListEndpoint(s)

	// применяем промежуточное ПО
	for _, m := range middlewares {
		eps.GetEmployeesSelectList = m(eps.GetEmployeesSelectList)
		eps.GetPositionsSelectList = m(eps.GetPositionsSelectList)
		eps.GetLevelAccessesSelectList = m(eps.GetLevelAccessesSelectList)
		eps.GetFreedomTypesSelectList = m(eps.GetFreedomTypesSelectList)
	}
}

func makeEmployeesSelectListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		domains, err := s.GetSelectListEmployees(ctx)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.SelectList, len(domains))
		for i := range domains {
			dtoRes[i] = toSelectListsDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}

func makePositionsSelectListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		domains, err := s.GetSelectListPosition(ctx)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.SelectList, len(domains))
		for i := range domains {
			dtoRes[i] = toSelectListsDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}

func makeLevelAccessesSelectListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		domains, err := s.GetSelectListLevelAccesses(ctx)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.LevelAccess, len(domains))
		for i := range domains {
			dtoRes[i] = toLevelAccessDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}

func makeFreedomTypesSelectListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		domains, err := s.GetSelectListFreedomType(ctx)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.SelectList, len(domains))
		for i := range domains {
			dtoRes[i] = toSelectListsDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}
