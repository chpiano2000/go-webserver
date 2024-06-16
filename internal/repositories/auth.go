package repositories

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"

	"github.com/go-webserver/internal/interfaces/auth"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authRepo struct{}

func NewAuthRepo() auth.AuthAdapter {
	return &authRepo{}
}

func (au *authRepo) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (au *authRepo) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (au *authRepo) GenerateKeyPair() ([]byte, []byte, error) {
	// Generate a new RSA private key with 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Encode the private key to the PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privatePEM := pem.EncodeToMemory(privateKeyPEM)

	// Extract the public key from the private key
	publicKey := &privateKey.PublicKey

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}
	publicPEM := pem.EncodeToMemory(publicKeyPEM)

	return privatePEM, publicPEM, nil
}

func (au *authRepo) GenerateTokens(accountID string, email string, privateKey []byte) (string, string, error) {
	// Create a new JWT token with claims
	accessClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        email,                            // Subject (user identifier)
		"iss":        "go-webserver",                   // Issuer
		"exp":        time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":        time.Now().Unix(),                // Issued at
		"account_id": accountID,
	})

	refreshClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":        email,                            // Subject (user identifier)
		"iss":        "go-webserver",                   // Issuer
		"exp":        time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat":        time.Now().Unix(),                // Issued at
		"account_id": accountID,
	})

	accessString, err := accessClaims.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}

	refreshString, err := refreshClaims.SignedString(privateKey)
	if err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}
