package api

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-gonic/gin"
	"github.com/tomaszczerminski/test/pkg/api/endpoints"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(helmet.Default())
	v1 := r.Group("/v1")
	{
		v1.GET("/ports", endpoints.GetPorts())
	}
	return r
}
