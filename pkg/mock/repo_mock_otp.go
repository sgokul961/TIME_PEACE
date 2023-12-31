// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/otp.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "gokul.go/pkg/utils/models"
)

// MockOtpRepository is a mock of OtpRepository interface.
type MockOtpRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOtpRepositoryMockRecorder
}

// MockOtpRepositoryMockRecorder is the mock recorder for MockOtpRepository.
type MockOtpRepositoryMockRecorder struct {
	mock *MockOtpRepository
}

// NewMockOtpRepository creates a new mock instance.
func NewMockOtpRepository(ctrl *gomock.Controller) *MockOtpRepository {
	mock := &MockOtpRepository{ctrl: ctrl}
	mock.recorder = &MockOtpRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOtpRepository) EXPECT() *MockOtpRepositoryMockRecorder {
	return m.recorder
}

// FindUserByMobileNumber mocks base method.
func (m *MockOtpRepository) FindUserByMobileNumber(phone string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByMobileNumber", phone)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FindUserByMobileNumber indicates an expected call of FindUserByMobileNumber.
func (mr *MockOtpRepositoryMockRecorder) FindUserByMobileNumber(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByMobileNumber", reflect.TypeOf((*MockOtpRepository)(nil).FindUserByMobileNumber), phone)
}

// UserDeatilsUsingPhone mocks base method.
func (m *MockOtpRepository) UserDeatilsUsingPhone(phone string) (models.UserDeatilsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserDeatilsUsingPhone", phone)
	ret0, _ := ret[0].(models.UserDeatilsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserDeatilsUsingPhone indicates an expected call of UserDeatilsUsingPhone.
func (mr *MockOtpRepositoryMockRecorder) UserDeatilsUsingPhone(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserDeatilsUsingPhone", reflect.TypeOf((*MockOtpRepository)(nil).UserDeatilsUsingPhone), phone)
}
