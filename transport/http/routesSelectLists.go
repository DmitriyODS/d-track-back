package http

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "gitlab.com/ddda/d-track/d-track-back/endpoints/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
	"net/http"
)

func initSelectListsRoutes(r *gin.RouterGroup, svcEps v1.Endpoints) {
	r.GET("/positions", func(c *gin.Context) {
		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetPositionsSelectList(newCtx, nil)
		if err != nil {
			log.Println("GetPositionsList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/employees", func(c *gin.Context) {
		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetEmployeesSelectList(newCtx, nil)
		if err != nil {
			log.Println("GetEmployeesList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/freedomTypes", func(c *gin.Context) {
		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetFreedomTypesSelectList(newCtx, nil)
		if err != nil {
			log.Println("GetFreedomTypesList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/levelAccesses", func(c *gin.Context) {
		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetLevelAccessesSelectList(newCtx, nil)
		if err != nil {
			log.Println("GetLevelAccessesList err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}
