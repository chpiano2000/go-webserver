package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-webserver/pkg/utils"

	log "github.com/sirupsen/logrus"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				message := utils.InternalServerError
				if castedErr, ok := err.(utils.RecipeMessage); ok {
					message = castedErr
				}
				resp := utils.Serialize(c, message)
				log.Infof("Response: %#v", resp)
				c.JSON(resp.Status, resp)
			}
		}()
		c.Next()
	}
}
