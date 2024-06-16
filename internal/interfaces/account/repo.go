package account

import "github.com/go-webserver/internal/models"

type AccountRepo interface {
	GetByEmail(email string) (*models.Account, error)
	InsertAccount(email string, name string, passwordHashed string) (string, error)
}
