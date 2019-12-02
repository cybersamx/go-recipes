package api

import (
	"math/rand"
	"strings"
	"time"
)

// AccountService represents a service that provides functions for managing user accounts
// in an application.
type AccountService struct {
	accountModel AccountModel
}

// NewAccountService creates an instance of AccountService.
func NewAccountService(accountModel AccountModel) AccountService {
	return AccountService{
		accountModel: accountModel,
	}
}

// ForgotPassword resets the password of the user associated with the passed email by
// replacing the existing password with a randomly generated password.
func (as *AccountService) ForgotPassword(email string) (string, error) {
	pwd := newRandomPassword()

	err := as.accountModel.UpdateAccount(email, pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}

func newRandomPassword() string {
	pool := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-.")
	rand.Seed(time.Now().Unix())
	var buffer strings.Builder
	for i := 0; i < 8; i++ {
		buffer.WriteRune(pool[rand.Intn(len(pool))])
	}

	return buffer.String()
}
