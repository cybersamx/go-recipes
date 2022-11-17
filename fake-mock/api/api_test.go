package api

import (
	"errors"
	"testing"

	"github.com/cybersamx/go-recipes/fake-mock/api/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAPIUsingFake(t *testing.T) {
	// Setup
	email := "sam@example.com"
	fam := mocks.NewFakeAccountModel()
	service := NewAccountService(fam)
	fam.AddAccount("sam@example.com", "12345678")
	fam.AddAccount("linda@example.com", "abcdefgh")

	// Run
	pwd, err := service.ForgotPassword(email)

	// Validation
	assert.NoError(t, err)
	assert.NotEmpty(t, pwd)

	foundAcct, err := fam.GetAccount(email)
	assert.NoError(t, err)
	assert.NotNil(t, foundAcct)
	assert.Equal(t, pwd, foundAcct.Password)
}

func TestAPIUsingMock(t *testing.T) {
	// Setup
	email := "sam@example.com"
	mam := mocks.NewTestifyMockAccountModel()
	service := NewAccountService(mam)
	mam.On("UpdateAccount", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	mam.On("AddAccount", mock.Anything, mock.Anything).Return(func(e, pwd string) error {
		// The following is strictly for demo. It doesn't make logic sense.
		if e == email {
			return nil
		}

		return errors.New("mismatched password")
	})

	// We can call the mocks directly. For this, instead of passing a literal value to
	// Return(), we can pass a function. To do so, we need to change the mocked AddAccount.
	assert.NoError(t, mam.AddAccount(email, "my-password"))
	assert.Error(t, mam.AddAccount("wrong-email", "my-password"))

	// Run
	// Note: service.ForgotPassword is the actual code that we are unit testing.
	pwd, err := service.ForgotPassword(email)

	// Validation
	assert.NoError(t, err)
	assert.NotEmpty(t, pwd)
	// Asset that the expectations were met
	mam.AssertExpectations(t)
}

func TestAPIUsingMockgen(t *testing.T) {
	// Setup
	email := "sam@example.com"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mam := mocks.NewMockAccountModel(ctrl)
	service := NewAccountService(mam)

	mam.EXPECT().
		UpdateAccount(email, gomock.Any()).
		Return(nil)

	// Run
	// Note: service.ForgotPassword is the actual code that we are unit testing.
	pwd, err := service.ForgotPassword(email)

	// Validation
	assert.NoError(t, err)
	assert.NotEmpty(t, pwd)
}
