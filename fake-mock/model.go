package api

import "errors"

type Account struct {
    Email string
    Password string
}

type AccountModel interface {
    AddAccount(email, password string) error
    UpdateAccount(email, password string) error
    GetAccount(email string) (*Account, error)
}

// This account model will be used in production and is associated with
// a SQL database.

type PostgresAccountModel struct {}

func (pam *PostgresAccountModel) AddAccount(email, password string) error {
    return errors.New("not implemented yet")
}

func (pam *PostgresAccountModel) UpdateAccount(email, password string) error {
    return errors.New("not implemented yet")
}

func (pam *PostgresAccountModel) GetAccount(email string) (*Account, error) {
    return nil, errors.New("not implemented yet")
}
