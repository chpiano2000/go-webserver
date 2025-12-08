package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/go-webserver/internal/response"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("PANIC RECOVERED: %v", err)

				log.Errorf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

				resp := response.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Code:    "InternalServerError",
					Message: "Unexpected error occurred",
					Data:    nil,
				}

				c.JSON(http.StatusInternalServerError, resp)
				c.Abort()
			}
		}()
		c.Next()
	}
}
