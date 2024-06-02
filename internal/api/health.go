package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-webserver/internal/lib"
)

type HealthRouter struct {
	handler lib.RequestHandler
}

func (rr HealthRouter) Setup(group string) {
	api := rr.handler.Gin.Group(group)
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func NewHealthRouter(handler lib.RequestHandler) HealthRouter {
	healthRouter := HealthRouter{
		handler: handler,
	}
	return healthRouter
}
