package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webserver/internal/interfaces/account"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/internal/response"
	"github.com/go-webserver/internal/schemas"
	"github.com/go-webserver/pkg/utils"
)

type AccountController struct {
	service account.AccountUseCase
}

func NewAccountController(useCase account.AccountUseCase) AccountController {
	return AccountController{
		service: useCase,
	}
}

func (ac AccountController) Signup(c *gin.Context) {
	var signupSchemas schemas.AccountSchemaRequest
	if err := c.ShouldBindJSON(&signupSchemas); err != nil {
		resp := utils.Serialize(c, utils.UnprocessableEntity)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resp)
		return
	}

	signupRequest := models.SignupRequest{
		Name:     signupSchemas.Name,
		Email:    signupSchemas.Email,
		Password: signupSchemas.Password,
	}

	account := ac.service.Signup(&signupRequest)
	if account.Err != nil {
		panic(account.Err)
	}
	successCode := "AccountCreated"
	successMessage := "Account Created Successfully"
	c.JSON(http.StatusCreated, response.Created(successCode, successMessage, account.Account))
}
