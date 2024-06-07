package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecipeMessage struct {
	Message     string
	MessageCode string
	Status      int
	Data        map[string]string
}

func (rm RecipeMessage) Error() string {
	return fmt.Sprintf("RecipeMessage: %s", rm.Message)
}

func (rm RecipeMessage) IsError(err error) bool {
	msgErr, ok := err.(RecipeMessage)
	if !ok {
		return false
	}
	return msgErr.MessageCode == rm.MessageCode
}

type Response struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Status  int    `json:"status"`
}

func Serialize(c *gin.Context, message interface{}) Response {
	respMessage := InternalServerError
	switch messageType := message.(type) {
	case RecipeMessage:
		respMessage = message.(RecipeMessage)
		return Response{
			Message: respMessage.Message,
			Code:    respMessage.MessageCode,
			Status:  respMessage.Status,
		}
	case error:
		return Response{
			Message: respMessage.Message,
			Code:    respMessage.MessageCode,
			Status:  respMessage.Status,
		}
	default:
		panic(fmt.Sprintf("Unknown message type: %T", messageType))
	}
}

var (
	OK = RecipeMessage{
		Message:     "Ok",
		MessageCode: "OK",
		Status:      http.StatusOK,
	}
	InternalServerError = RecipeMessage{
		Message:     "Internal server error",
		MessageCode: "InternalServerError",
		Status:      http.StatusInternalServerError,
	}
	UnprocessableEntity = RecipeMessage{
		Message:     "Unprocessable entity",
		MessageCode: "UnprocessableEntity",
		Status:      http.StatusUnprocessableEntity,
	}
)
