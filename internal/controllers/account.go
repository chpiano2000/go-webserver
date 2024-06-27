package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webserver/constants"
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

// Signup godoc
// @Summary Signup
// @Description User Signup
// @Tags Account
// @Accept json
// @Produce json
// @Param payload body schemas.SignupSchemaRequest true "user signup payload"
// @Success     200         {object}    models.SignupResponse
// @Failure     400         {object}    response.ErrorResponse
// @Failure     422         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /signup [post]
func (ac AccountController) Signup(c *gin.Context) {
	var signupSchemas schemas.SignupSchemaRequest
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

// Login godoc
// @Summary Login
// @Description User Login
// @Tags Account
// @Accept json
// @Produce json
// @Param payload body schemas.LoginSchemaRequest true "user login payload"
// @Success     200         {object}    models.LoginResponse
// @Failure     400         {object}    response.ErrorResponse
// @Failure     422         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /login [post]
func (ac AccountController) Login(c *gin.Context) {
	var loginSchemas schemas.LoginSchemaRequest
	if err := c.ShouldBindJSON(&loginSchemas); err != nil {
		resp := utils.Serialize(c, utils.UnprocessableEntity)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, resp)
		return
	}

	loginRequest := models.LoginRequest{
		Email:    loginSchemas.Email,
		Password: loginSchemas.Password,
	}

	account := ac.service.Login(&loginRequest)
	if account.Err != nil {
		panic(account.Err)
	}
	successCode := "LoginCreated"
	successMessage := "Login Successfully"
	c.JSON(http.StatusCreated, response.Created(successCode, successMessage, account.Account))
}

// Logout godoc
// @Summary Logout
// @Description User Logout
// @Tags Account
// @Accept json
// @Produce json
// @Success     200         {object}    models.SignupResponse
// @Failure     400         {object}    response.ErrorResponse
// @Failure     422         {object}    response.ErrorResponse
// @Failure     500         {object}    response.ErrorResponse
// @Router /logout [post]
func (ac AccountController) Logout(c *gin.Context) {
	accountId := c.Request.Header[constants.ClientHeader][0]
	err := ac.service.Logout(accountId)
	if err != nil {
		panic(err)
	}
	successCode := "LogoutSuccess"
	successMessage := "Logout Successfully"
	c.JSON(http.StatusCreated, response.Deleted(successCode, successMessage, nil))
}
