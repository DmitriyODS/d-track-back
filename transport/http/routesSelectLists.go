package http

import (
	"context"
	"github.com/gin-gonic/gin"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	v1 "gitlab.com/ddda/d-track/d-track-back/endpoints/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
	"net/http"
)

func initSelectListsRoutes(r *gin.RouterGroup, svcEps v1.Endpoints) {
	r.GET("/positions", func(c *gin.Context) {
		list := dtoV1.RequestList{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&list); err != nil {
			c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
			return
		}

		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetPositionsSelectList(newCtx, list)
		if err != nil {
			log.Println("GetPositionsList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/employees", func(c *gin.Context) {
		list := dtoV1.RequestList{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&list); err != nil {
			c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
			return
		}

		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetEmployeesSelectList(newCtx, list)
		if err != nil {
			log.Println("GetEmployeesList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/freedomTypes", func(c *gin.Context) {
		list := dtoV1.RequestList{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&list); err != nil {
			c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
			return
		}

		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetFreedomTypesSelectList(newCtx, list)
		if err != nil {
			log.Println("GetFreedomTypesList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/levelAccesses", func(c *gin.Context) {
		list := dtoV1.RequestList{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&list); err != nil {
			c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
			return
		}

		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetLevelAccessesSelectList(newCtx, list)
		if err != nil {
			log.Println("GetLevelAccessesList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}
