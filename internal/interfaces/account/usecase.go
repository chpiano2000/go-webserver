package account

import "github.com/go-webserver/internal/models"

type AccountUseCase interface {
	Signup(user *models.SignupRequest) (resp models.SignupResponse)
}
