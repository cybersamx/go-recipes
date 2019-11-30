// The account model that is actually used in production will be connected to a
// database like Postgres or MySQL. This mocked account model records the
// activities of the how callers interface with the mocked account model and
// returns what the values that the object is instructed.

package api

import (
    "github.com/stretchr/testify/mock"
)

type MockAccountModel struct {
    mock.Mock
}

func NewMockAccountModel() *MockAccountModel {
    return &MockAccountModel{}
}

func (mam *MockAccountModel) AddAccount(email, password string) error {
    args := mam.Called(email, password)

    return args.Error(0)
}

func (mam *MockAccountModel) UpdateAccount(email, password string) error {
    args := mam.Called(email, password)

    return args.Error(0)
}

func (mam *MockAccountModel) GetAccount(email string) (*Account, error) {
    args := mam.Called(email)

    return args.Get(0).(*Account), args.Error(1)
}
