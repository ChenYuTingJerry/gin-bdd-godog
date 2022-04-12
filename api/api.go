package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ChenYuTingJerry/gin-bdd-godog/api/health"
	"github.com/ChenYuTingJerry/gin-bdd-godog/api/version"
)

type Response struct {
	Version string `json:"version"`
}

func SetRouter() *gin.Engine {
	r := gin.Default()
	routeGroup := r.Group("/")
	health.New(routeGroup)
	version.New(routeGroup)
	return r
}
