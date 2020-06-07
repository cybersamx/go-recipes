package model

// Account represents the user account entity.
type Account struct {
	Email    string
	Password string
}

// AccountModel is an abstract interface for access and saving accounts.
type AccountModel interface {
	AddAccount(email, password string) error
	UpdateAccount(email, password string) error
	GetAccount(email string) (*Account, error)
}
