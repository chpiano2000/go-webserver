package auth

import "crypto/rsa"

type AuthAdapter interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error)
	ConvertRSAToString(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, string)
	GenerateTokens(accountID string, email string, privateKey *rsa.PrivateKey) (string, string, error)
}
