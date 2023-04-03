package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "gitlab.com/ddda/d-track/d-track-back/endpoints/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"net/http"
	"time"
)

func NewHttpTransport(svcEps v1.Endpoints) http.Handler {
	r := gin.New()

	// задаём обработчик для "неизвестного" пути
	r.NoRoute(noRouteHandler)

	// задаём обработчик для "неизвестного" метода
	r.NoMethod(noMethodHandler)

	// создаём конфиг для CORS
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{global.HeaderAuthenticationKey, "Content-Type"},
		MaxAge:          12 * time.Hour,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	}

	// дрбалвяем ещё много всякого разного (middlewares)
	r.Use(gin.CustomRecovery(RecoveryPanic), gin.Logger(), cors.New(corsConfig))

	// API version 1
	apiV1 := r.Group("/v1")

	initAuthRoutes(apiV1.Group("/auth"), svcEps)
	initEmployeeRoutes(apiV1.Group("/employees", JWTAuth()), svcEps)
	initSelectListsRoutes(apiV1.Group("/lists", JWTAuth()), svcEps)

	return r
}
