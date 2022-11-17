// The account model that is actually used in production will be connected to a
// database like Postgres or MySQL. This mocked account model records the
// activities of the how callers interface with the mocked account model and
// returns what the values that the object is instructed.

package mocks

import (
	"github.com/cybersamx/go-recipes/fake-mock/model"
	"github.com/stretchr/testify/mock"
)

// TestifyMockAccountModel implements the AccountModel by mocking the data access layer.
type TestifyMockAccountModel struct {
	mock.Mock
}

// NewTestifyMockAccountModel creates an instance of TestifyMockAccountModel.
func NewTestifyMockAccountModel() *TestifyMockAccountModel {
	return &TestifyMockAccountModel{}
}

// AddAccount sets up the function to be stubbed with actual outputs to mimic
// adding a new account to the database.
func (mam *TestifyMockAccountModel) AddAccount(email, password string) error {
	ret := mam.Called(email, password)

	var r error
	arg := ret.Get(0)
	if rf, ok := arg.(func(string, string) error); ok {
		r = rf(email, password)
	} else {
		if arg == nil {
			return nil
		}

		r = arg.(error) // We can also use ret.Error(0)
	}

	return r
}

// UpdateAccount sets up the function to be stubbed with actual outputs to mimic
// updating an account in the database with passed email and password.
func (mam *TestifyMockAccountModel) UpdateAccount(email, password string) error {
	ret := mam.Called(email, password)

	return ret.Error(0)
}

// GetAccount sets up the function to be stubbed with actual outputs to mimic
// retrieving an account associated with the passed email.
func (mam *TestifyMockAccountModel) GetAccount(email string) (*model.Account, error) {
	ret := mam.Called(email)

	return ret.Get(0).(*model.Account), ret.Error(1)
}
