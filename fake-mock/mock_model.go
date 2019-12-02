// The account model that is actually used in production will be connected to a
// database like Postgres or MySQL. This mocked account model records the
// activities of the how callers interface with the mocked account model and
// returns what the values that the object is instructed.

package api

import (
	"github.com/stretchr/testify/mock"
)

// MockAccountModel implements the AccountModel by mocking the data access layer.
type MockAccountModel struct {
	mock.Mock
}

// NewMockAccountModel creates an instance of MockAccountModel.
func NewMockAccountModel() *MockAccountModel {
	return &MockAccountModel{}
}

// AddAccount sets up the function to be stubbed with actual outputs to mimic
// adding a new account to the database.
func (mam *MockAccountModel) AddAccount(email, password string) error {
	args := mam.Called(email, password)

	return args.Error(0)
}

// UpdateAccount sets up the function to be stubbed with actual outputs to mimic
// updating an account in the database with passed email and password.
func (mam *MockAccountModel) UpdateAccount(email, password string) error {
	args := mam.Called(email, password)

	return args.Error(0)
}

// GetAccount sets up the function to be stubbed with actual outputs to mimic
// retrieving an account associated with the passed email.
func (mam *MockAccountModel) GetAccount(email string) (*Account, error) {
	args := mam.Called(email)

	return args.Get(0).(*Account), args.Error(1)
}
