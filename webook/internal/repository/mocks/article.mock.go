// Code generated by MockGen. DO NOT EDIT.
// Source: ./webook/internal/repository/article.go
//
// Generated by this command:
//
//	mockgen -source=./webook/internal/repository/article.go -package=repomocks -destination=./webook/internal/repository/mocks/article.mock.go
//

// Package repomocks is a generated GoMock package.
package repomocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/ischeng28/basic-go/webook/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleRepository is a mock of ArticleRepository interface.
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
	isgomock struct{}
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository.
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance.
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleRepository) Create(ctx context.Context, art domain.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, art)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArticleRepositoryMockRecorder) Create(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleRepository)(nil).Create), ctx, art)
}

// Sync mocks base method.
func (m *MockArticleRepository) Sync(ctx context.Context, art domain.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", ctx, art)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockArticleRepositoryMockRecorder) Sync(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockArticleRepository)(nil).Sync), ctx, art)
}

// Update mocks base method.
func (m *MockArticleRepository) Update(ctx context.Context, art domain.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, art)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleRepositoryMockRecorder) Update(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleRepository)(nil).Update), ctx, art)
}