// Code generated by MockGen. DO NOT EDIT.
// Source: ./webook/internal/repository/code.go
//
// Generated by this command:
//
//	mockgen -source=./webook/internal/repository/code.go -package=repomocks -destination=./webook/internal/repository/mocks/code.mock.go
//

// Package repomocks is a generated GoMock package.
package repomocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCodeRepository is a mock of CodeRepository interface.
type MockCodeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCodeRepositoryMockRecorder
	isgomock struct{}
}

// MockCodeRepositoryMockRecorder is the mock recorder for MockCodeRepository.
type MockCodeRepositoryMockRecorder struct {
	mock *MockCodeRepository
}

// NewMockCodeRepository creates a new mock instance.
func NewMockCodeRepository(ctrl *gomock.Controller) *MockCodeRepository {
	mock := &MockCodeRepository{ctrl: ctrl}
	mock.recorder = &MockCodeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeRepository) EXPECT() *MockCodeRepositoryMockRecorder {
	return m.recorder
}

// Set mocks base method.
func (m *MockCodeRepository) Set(ctx context.Context, biz, phone, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, biz, phone, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCodeRepositoryMockRecorder) Set(ctx, biz, phone, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCodeRepository)(nil).Set), ctx, biz, phone, code)
}

// Verify mocks base method.
func (m *MockCodeRepository) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, biz, phone, code)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockCodeRepositoryMockRecorder) Verify(ctx, biz, phone, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockCodeRepository)(nil).Verify), ctx, biz, phone, code)
}