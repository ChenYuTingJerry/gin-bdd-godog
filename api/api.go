package api

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Version string `json:"version"`
}

func getVersion(c *gin.Context) {
	res := Response{
		Version: "v2.12.3",
	}
	c.JSON(200, res)
}

func getHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health", getHealth)
	r.GET("/version", getVersion)
	return r
}
