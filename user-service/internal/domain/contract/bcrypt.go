package contract

type IBcryptAdapter interface {
	Hash(password string, salt int) (string, error)
}
