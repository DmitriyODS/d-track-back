package v1

import (
	"context"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// TaskEndpoints - конечные точки для работы с заявками
type TaskEndpoints struct {
	GetTasksList global.Endpoint
	GetTaskByID  global.Endpoint
	TaskStore    global.Endpoint
}

func (eps *TaskEndpoints) makeTaskEndpoints(s service.Service, middlewares ...global.Middleware) {
	eps.GetTasksList = makeGetTasksListEndpoint(s)
	eps.GetTaskByID = makeGetTaskByIDEndpoint(s)
	eps.TaskStore = makeTaskStoreEndpoint(s)

	// применяем промежуточное ПО
	for _, m := range middlewares {
		eps.GetTasksList = m(eps.GetTasksList)
		eps.GetTaskByID = m(eps.GetTaskByID)
		eps.TaskStore = m(eps.TaskStore)
	}
}

func makeGetTasksListEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestTaskListFilters)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		domains, err := s.GetListTasks(ctx, req.NumberFilter, req.IsArchive)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		dtoRes := make([]dtoV1.Task, len(domains))
		for i := range domains {
			dtoRes[i] = toTaskDTO(domains[i])
		}

		return global.NewResponseData(dtoRes), nil
	}
}

func makeGetTaskByIDEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.RequestByID)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		task, err := s.GetTaskByID(ctx, req.ID)
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(toTaskDTO(task)), nil
	}
}

func makeTaskStoreEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.Task)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		id, err := s.StoreTask(ctx, fromTaskDTO(req))
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(id), nil
	}
}
