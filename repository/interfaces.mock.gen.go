// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetProfile mocks base method.
func (m *MockRepositoryInterface) GetProfile(ctx context.Context, input GetProfileInput) (*GetProfileOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx, input)
	ret0, _ := ret[0].(*GetProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockRepositoryInterfaceMockRecorder) GetProfile(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).GetProfile), ctx, input)
}

// GetUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) GetUserByPhoneNumber(ctx context.Context, input GetUserByPhoneNumberInput) (*GetUserByPhoneNumberOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, input)
	ret0, _ := ret[0].(*GetUserByPhoneNumberOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhoneNumber(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhoneNumber), ctx, input)
}

// SaveUser mocks base method.
func (m *MockRepositoryInterface) SaveUser(ctx context.Context, input SaveUserInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockRepositoryInterfaceMockRecorder) SaveUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockRepositoryInterface)(nil).SaveUser), ctx, input)
}

// UpdateUserLoginCount mocks base method.
func (m *MockRepositoryInterface) UpdateUserLoginCount(ctx context.Context, input UpdateUserCountInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserLoginCount", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserLoginCount indicates an expected call of UpdateUserLoginCount.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateUserLoginCount(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserLoginCount", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateUserLoginCount), ctx, input)
}

// UpdateUserProfile mocks base method.
func (m *MockRepositoryInterface) UpdateUserProfile(ctx context.Context, input UpdateUserProfileInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProfile", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserProfile indicates an expected call of UpdateUserProfile.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateUserProfile(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProfile", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateUserProfile), ctx, input)
}
