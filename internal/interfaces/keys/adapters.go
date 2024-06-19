package keys

// import "github.com/go-webserver/internal/models"

type KeyAdapter interface {
	InsertKeys(accountId string, publicKey string, privateKey string, refreshToken string) error
	// GetKeyByAccount(accountId string) (*models)
}
