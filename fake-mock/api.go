package api

import (
	"math/rand"
	"strings"
	"time"
)

type AccountService struct {
	accountModel AccountModel
}

func NewAccountService(accountModel AccountModel) AccountService {
	return AccountService{
		accountModel: accountModel,
	}
}

func (as *AccountService) ForgotPassword(email string) (string, error) {
	pwd := NewRandomPassword()

	err := as.accountModel.UpdateAccount(email, pwd)
	if err != nil {
		return "", err
	}

	return pwd, nil
}

func NewRandomPassword() string {
	pool := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-.")
	rand.Seed(time.Now().Unix())
	var buffer strings.Builder
	for i := 0; i < 8; i++ {
		buffer.WriteRune(pool[rand.Intn(len(pool))])
	}

	return buffer.String()
}
