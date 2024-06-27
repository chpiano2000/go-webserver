package keys

type KeyAdapter interface {
	InsertKeys(accountId string, publicKey string, privateKey string, refreshToken string) error
	RemoveKeysByID(accountId string) error
}
