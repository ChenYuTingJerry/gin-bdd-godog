package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type impl struct {
}

func New(router *gin.RouterGroup) impl {
	im := impl{}
	router.GET("/version", im.getVersion)
	return im
}

func (im impl) getVersion(c *gin.Context) {
	res := struct {
		Version string `json:"version"`
	}{Version: "v1.0.0"}

	c.JSON(http.StatusOK, res)
}
