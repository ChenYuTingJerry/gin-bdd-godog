package echo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type impl struct {
}

type Request struct {
	Echo string `json:"echo"`
}

func New(router *gin.RouterGroup) impl {
	im := impl{}
	router.POST("/echo", im.echo)
	return im
}

func (im impl) echo(c *gin.Context) {
	reqBody := Request{}
	c.BindJSON(&reqBody)

	c.JSON(http.StatusOK, reqBody)
}
