package auth

type AuthAdapter interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	GenerateKeyPair() ([]byte, []byte, error)
	GenerateTokens(accountID string, email string, privateKey []byte) (string, string, error)
}

