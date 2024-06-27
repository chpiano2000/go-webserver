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
		Id:           accountID,
		Name:         user.Name,
		Email:        user.Email,
		AccessToken:  "",
		RefreshToken: "",
	}
	return
}

func (ac *accountService) Login(user *models.LoginRequest) (resp models.LoginResponse) {
	// Check if an account has regitered
	userInDB, err := ac.accountRepo.GetByEmail(user.Email)
	if err != nil {
		if err.Error() == utils.AccountNotFound.Error() {
			logger.Errorf("accountService::Signup::GetByEmail %v", err)
			resp.Err = utils.AccountExists
			return
		} else {
			logger.Errorf("accountService::Signup::GetByEmail %v", err)
			resp.Err = err
			return
		}
	}

	// Check if password is correct
	passwordMatched := ac.authRepo.CheckPasswordHash(user.Password, userInDB.PasswordHashed)
	if !passwordMatched {
		logger.Errorf("accountService::Signup::GetByEmail %v", err)
		resp.Err = utils.AccountExists
		return
	}

	privateKey, publicKey, err := ac.authRepo.GenerateKeyPair()
	if err != nil {
		logger.Errorf("accountService::Signup::GenerateKeyPair %v", err)
		resp.Err = err
		return
	}

	// Generate token pair
	accessToken, refreshToken, err := ac.authRepo.GenerateTokens(userInDB.Id, user.Email, privateKey)
	if err != nil {
		logger.Errorf("accountService::Signup::GenerateTokens %v", err)
		resp.Err = err
		return
	}

	privateKeyString, publicKeyString := ac.authRepo.ConvertRSAToString(privateKey, publicKey)

	err = ac.keysRepo.InsertKeys(userInDB.Id, publicKeyString, privateKeyString, refreshToken)
	if err != nil {
		logger.Errorf("accountService::Signup::InsertKeys %v", err)
		resp.Err = err
		return
	}
	resp.Account = &models.AccountAuthResponse{
		Id:           userInDB.Id,
		Name:         userInDB.Name,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return
}

func (ac *accountService) Logout(keyId string) error {
	err := ac.keysRepo.RemoveKeysByID(keyId)
	if err != nil {
		logger.Errorf("accountService::Logout::RemoveKeysByID %v", err)
		return err
	}

	return nil
}
