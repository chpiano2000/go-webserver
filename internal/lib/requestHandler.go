package lib

import (
	"github.com/gin-gonic/gin"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates a new request handler
func NewRequestHandler() RequestHandler {
	// gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.New()
	engine.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/pathsNotToLog/"),
		gin.Recovery(),
	)
	return RequestHandler{Gin: engine}
}
