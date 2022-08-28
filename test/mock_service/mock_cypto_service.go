// Code generated by MockGen. DO NOT EDIT.
// Source: service/cypto_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCyptoServiceRepository is a mock of CyptoServiceRepository interface.
type MockCyptoServiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCyptoServiceRepositoryMockRecorder
}

// MockCyptoServiceRepositoryMockRecorder is the mock recorder for MockCyptoServiceRepository.
type MockCyptoServiceRepositoryMockRecorder struct {
	mock *MockCyptoServiceRepository
}

// NewMockCyptoServiceRepository creates a new mock instance.
func NewMockCyptoServiceRepository(ctrl *gomock.Controller) *MockCyptoServiceRepository {
	mock := &MockCyptoServiceRepository{ctrl: ctrl}
	mock.recorder = &MockCyptoServiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCyptoServiceRepository) EXPECT() *MockCyptoServiceRepositoryMockRecorder {
	return m.recorder
}

// ComparePasswords mocks base method.
func (m *MockCyptoServiceRepository) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePasswords", hashedPwd, plainPwd)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ComparePasswords indicates an expected call of ComparePasswords.
func (mr *MockCyptoServiceRepositoryMockRecorder) ComparePasswords(hashedPwd, plainPwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePasswords", reflect.TypeOf((*MockCyptoServiceRepository)(nil).ComparePasswords), hashedPwd, plainPwd)
}

// HashAndSalt mocks base method.
func (m *MockCyptoServiceRepository) HashAndSalt(pwd []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashAndSalt", pwd)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashAndSalt indicates an expected call of HashAndSalt.
func (mr *MockCyptoServiceRepositoryMockRecorder) HashAndSalt(pwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashAndSalt", reflect.TypeOf((*MockCyptoServiceRepository)(nil).HashAndSalt), pwd)
}
