// The account model that is actually used in production will be connected to a
// database like Postgres or MySQL. This fake account model mimics data access
// to the database using in-code map data type.

package api

import "fmt"

type FakeAccountModel struct {
    accounts map[string]*Account
}

func NewFakeAccountModel() *FakeAccountModel {
    return &FakeAccountModel{
        accounts: make(map[string]*Account),
    }
}

func (fam *FakeAccountModel) AddAccount(email, password string) error {
    account := Account{
        Email:    email,
        Password: password,
    }

    fam.accounts[email] = &account

    return nil
}

func (fam *FakeAccountModel) UpdateAccount(email, password string) error {
    foundAcct := fam.accounts[email]
    if foundAcct == nil {
        return fmt.Errorf("can't find account with eamil %s", email)
    }

    foundAcct.Password = password

    return nil
}

func (fam *FakeAccountModel) GetAccount(email string) (*Account, error) {
    foundAcct := fam.accounts[email]
    if foundAcct == nil {
        return nil, fmt.Errorf("can't find account with eamil %s", email)
    }

    return foundAcct, nil
}

