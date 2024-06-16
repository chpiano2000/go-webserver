package models

import "time"

type Account struct {
	Id             string
	Name           string
	Email          string
	PasswordHashed string
	Status         string
	Verify         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type AccountAuthResponse struct {
	Id          string
	Name        string
	Email       string
	AccessToken string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email    string
	Name     string
	Password string
}

type SignupResponse struct {
	Account *AccountAuthResponse
	Err     error
}
