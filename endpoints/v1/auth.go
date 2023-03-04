package v1

import (
	"context"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/service"
)

// AuthEndpoints - конечные точки для аутентификации в сервисе
type AuthEndpoints struct {
	AuthenticationByLogin global.Endpoint
}

func (eps *AuthEndpoints) makeAuthEndpoints(s service.Service, middlewares ...global.Middleware) {
	eps.AuthenticationByLogin = makeAuthenticationByLoginEndpoint(s)

	// применяем промежуточное ПО
	for _, m := range middlewares {
		eps.AuthenticationByLogin = m(eps.AuthenticationByLogin)
	}
}

func makeAuthenticationByLoginEndpoint(s service.Service) global.Endpoint {
	return func(ctx context.Context, request interface{}) (response global.ResponseData, err error) {
		req, ok := request.(dtoV1.Auth)
		if !ok {
			return global.NewErrResponseData(global.IncorrectBodyRequestErr), global.IncorrectBodyRequestErr
		}

		auth, err := s.AuthenticationByLogin(ctx, fromAuthDTO(req))
		if err != nil {
			return global.NewErrResponseData(err), err
		}

		return global.NewResponseData(toAuthDTO(auth)), nil
	}
}
