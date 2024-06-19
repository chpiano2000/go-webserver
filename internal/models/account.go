package models

import "time"

type Account struct {
	Id             string    `bson:"_id" json:"id" `
	Name           string    `bson:"name" json:"name"`
	Email          string    `bson:"email" json:"email"`
	PasswordHashed string    `bson:"passwordHashed" json:"passwordHashed"`
	Status         string    `bson:"status" json:"status"`
	Verify         bool      `bson:"verify" json:"verify"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
}

type AccountAuthResponse struct {
	Id          string `json:"id" `
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Email    string
	Password string
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

type LoginResponse struct {
	Account *AccountAuthResponse
	Err     error
}
