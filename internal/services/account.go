package services

import (
	"fmt"

	"github.com/go-webserver/internal/interfaces/account"
	"github.com/go-webserver/internal/interfaces/auth"
	"github.com/go-webserver/internal/interfaces/keys"
	"github.com/go-webserver/internal/models"
	"github.com/go-webserver/pkg/utils"

	logger "github.com/sirupsen/logrus"
)

type accountService struct {
	accountRepo account.AccountRepo
	authRepo    auth.AuthAdapter
	keysRepo    keys.KeyAdapter
}

func NewAccountService(accountRepo account.AccountRepo, authRepo auth.AuthAdapter, keysRepo keys.KeyAdapter) account.AccountUseCase {
	return &accountService{accountRepo: accountRepo, authRepo: authRepo, keysRepo: keysRepo}
}

func (ac *accountService) Signup(user *models.SignupRequest) (resp models.SignupResponse) {
	userInDB, err := ac.accountRepo.GetByEmail(user.Email)
	if err != nil && err.Error() != utils.AccountNotFound.Error() {
		if userInDB.Email == user.Email {
			logger.Errorf("accountService::Signup::GetByEmail %v", err)
			resp.Err = utils.AccountExists
			return
		} else {
			logger.Errorf("accountService::Signup::GetByEmail %v", err)
			resp.Err = err
			return
		}
	}
	
	if userInDB != nil {
		if userInDB.Email == user.Email {
			logger.Errorf("accountService::Signup::GetByEmail %v", err)
			resp.Err = utils.AccountExists
			return
		}
	}

	hasedPassword, err := ac.authRepo.HashPassword(user.Password)
	if err != nil {
		logger.Errorf("accountService::Signup::HashPassword %v", err)
		panic(fmt.Errorf("cannot has password: %v", err))
	}

	accountID, err := ac.accountRepo.InsertAccount(user.Email, user.Name, hasedPassword)
	if err != nil {
		logger.Errorf("accountService::Signup::InsertAccount %v", err)
		resp.Err = err
		return
	}

	resp.Account = &models.AccountAuthResponse{
		Id:          accountID,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: "",
	}
	return
}
