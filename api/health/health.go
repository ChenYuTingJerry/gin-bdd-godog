package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type impl struct {
}

func New(router *gin.RouterGroup) impl {
	im := impl{}
	router.GET("/health", im.getHealth)
	return im
}

func (im impl) getHealth(c *gin.Context) {
	res := struct {
		Status string `json:"status"`
	}{Status: "ok"}

	c.JSON(http.StatusOK, res)
}
