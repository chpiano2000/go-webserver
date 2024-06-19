package api

import (
	"github.com/go-webserver/internal/controllers"
	"github.com/go-webserver/internal/lib"
)

type AccountRouter struct {
	accountController controllers.AccountController
	handler           lib.RequestHandler
}

func (rr AccountRouter) Setup(group string) {
	api := rr.handler.Gin.Group(group)
	{
		api.POST("/signup", rr.accountController.Signup)
		api.POST("/login", rr.accountController.Login)
	}
}

func NewAccountRouter(svc controllers.AccountController, handler lib.RequestHandler) AccountRouter {
	accountRouter := AccountRouter{
		accountController: svc,
		handler:           handler,
	}
	return accountRouter
}
