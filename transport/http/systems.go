package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"net/http"
)

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, global.NewErrResponseData(global.NotFoundErr))
}

func noMethodHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, global.NewErrResponseData(global.NotImplementedErr))
}
