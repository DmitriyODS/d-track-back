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

func initAuthRoutes(r *gin.RouterGroup, svcEps v1.Endpoints) {
	r.POST("/login", func(c *gin.Context) {
		data := dtoV1.Auth{}

		// выполняем привязку запроса к структуре запроса
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusInternalServerError, global.NewErrResponseData(global.InternalServerErr))
			return
		}

		newCtx := context.Background()

		// вызываем конечную точку обработки запроса
		resp, err := svcEps.AuthenticationByLogin(newCtx, data)
		if err != nil {
			log.Println("Login router err:", err)
			c.JSON(resp.CodeErr, resp)
			return
		}

		c.JSON(http.StatusOK, resp)
	})
}
