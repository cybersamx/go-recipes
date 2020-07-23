// Code generated by MockGen. DO NOT EDIT.
// Source: ./model/account.go

// Package mock_recipe is a generated GoMock package.
package mock_recipe

import (
	model "github.com/cybersamx/go-recipes/fake-mock/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountModel is a mock of AccountModel interface
type MockAccountModel struct {
	ctrl     *gomock.Controller
	recorder *MockAccountModelMockRecorder
}

// MockAccountModelMockRecorder is the mock recorder for MockAccountModel
type MockAccountModelMockRecorder struct {
	mock *MockAccountModel
}

// NewMockAccountModel creates a new mock instance
func NewMockAccountModel(ctrl *gomock.Controller) *MockAccountModel {
	mock := &MockAccountModel{ctrl: ctrl}
	mock.recorder = &MockAccountModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountModel) EXPECT() *MockAccountModelMockRecorder {
	return m.recorder
}

// AddAccount mocks base method
func (m *MockAccountModel) AddAccount(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccount", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAccount indicates an expected call of AddAccount
func (mr *MockAccountModelMockRecorder) AddAccount(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccount", reflect.TypeOf((*MockAccountModel)(nil).AddAccount), email, password)
}

// UpdateAccount mocks base method
func (m *MockAccountModel) UpdateAccount(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockAccountModelMockRecorder) UpdateAccount(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountModel)(nil).UpdateAccount), email, password)
}

// GetAccount mocks base method
func (m *MockAccountModel) GetAccount(email string) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", email)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockAccountModelMockRecorder) GetAccount(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountModel)(nil).GetAccount), email)
}