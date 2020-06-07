package model

import "errors"

// PostgresAccountModel represents the access layer to the SQL database.
type PostgresAccountModel struct{}

// AddAccount adds a new account to the SQL database.
func (pam *PostgresAccountModel) AddAccount(email, password string) error {
	return errors.New("not implemented yet")
}

// UpdateAccount updates the account in the SQL database with passed email and password.
func (pam *PostgresAccountModel) UpdateAccount(email, password string) error {
	return errors.New("not implemented yet")
}

// GetAccount retrieves the account in the SQL database associated with the passed email.
func (pam *PostgresAccountModel) GetAccount(email string) (*Account, error) {
	return nil, errors.New("not implemented yet")
}
