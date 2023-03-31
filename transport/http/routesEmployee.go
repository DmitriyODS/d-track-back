package http

import (
	"context"
	"github.com/gin-gonic/gin"
	dtoV1 "gitlab.com/ddda/d-track/d-track-back/dto/v1"
	v1 "gitlab.com/ddda/d-track/d-track-back/endpoints/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"log"
	"net/http"
	"strconv"
)

func initEmployeeRoutes(r *gin.RouterGroup, svcEps v1.Endpoints) {
	r.GET("/getList", func(c *gin.Context) {
		list := dtoV1.RequestList{}

		// выполняем привязку запроса к структуре запроса
		// TODO: пока этого не делаем, нужно разобраться с запросами GET и где отправлять фильтры
		//  в строке запроса???
		//if err := c.ShouldBindJSON(&list); err != nil {
		//	c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
		//	return
		//}

		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetEmployeesList(newCtx, list)
		if err != nil {
			log.Println("GetList router err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.GET("/:byID", func(c *gin.Context) {
		// берём параметр из строки запроса
		id, err := strconv.ParseUint(c.Param("byID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
			return
		}
		reqID := dtoV1.RequestByID{ID: id}

		// пытаемся получить мета-инфу о пользователе, чтобы сохранить её в новый контекст
		v, ok := c.Get(global.JwtClaimsCtxKey)
		if !ok {
			c.JSON(http.StatusUnauthorized, global.NewErrResponseData(global.UnauthorizedErr))
			return
		}

		newCtx := context.WithValue(context.Background(), global.JwtClaimsCtxKey, v)

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.GetEmployeeByID(newCtx, reqID)
		if err != nil {
			log.Println("EmployeeByID router err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.POST("/store", func(c *gin.Context) {
		data := dtoV1.Employee{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&data); err != nil {
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
		resp, err := svcEps.EmployeeStore(newCtx, data)
		if err != nil {
			log.Println("Store router err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.PUT("/store", func(c *gin.Context) {
		data := dtoV1.Employee{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&data); err != nil {
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
		resp, err := svcEps.EmployeeStore(newCtx, data)
		if err != nil {
			log.Println("Store router err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}
