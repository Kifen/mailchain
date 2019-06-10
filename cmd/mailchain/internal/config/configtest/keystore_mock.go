// Code generated by MockGen. DO NOT EDIT.
// Source: keystore.go

// Package configtest is a generated GoMock package.
package configtest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKeystoreSetter is a mock of KeystoreSetter interface
type MockKeystoreSetter struct {
	ctrl     *gomock.Controller
	recorder *MockKeystoreSetterMockRecorder
}

// MockKeystoreSetterMockRecorder is the mock recorder for MockKeystoreSetter
type MockKeystoreSetterMockRecorder struct {
	mock *MockKeystoreSetter
}

// NewMockKeystoreSetter creates a new mock instance
func NewMockKeystoreSetter(ctrl *gomock.Controller) *MockKeystoreSetter {
	mock := &MockKeystoreSetter{ctrl: ctrl}
	mock.recorder = &MockKeystoreSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeystoreSetter) EXPECT() *MockKeystoreSetterMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockKeystoreSetter) Set(keystoreType, keystorePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", keystoreType, keystorePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockKeystoreSetterMockRecorder) Set(keystoreType, keystorePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockKeystoreSetter)(nil).Set), keystoreType, keystorePath)
}
