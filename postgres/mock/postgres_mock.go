// Code generated by MockGen. DO NOT EDIT.
// Source: postgres.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	postgres "github.com/StasBoiko/test-work/postgres"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateAuthor mocks base method.
func (m *MockRepository) CreateAuthor(ctx context.Context, name string) (postgres.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAuthor", ctx, name)
	ret0, _ := ret[0].(postgres.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAuthor indicates an expected call of CreateAuthor.
func (mr *MockRepositoryMockRecorder) CreateAuthor(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAuthor", reflect.TypeOf((*MockRepository)(nil).CreateAuthor), ctx, name)
}

// CreateBook mocks base method.
func (m *MockRepository) CreateBook(ctx context.Context, title string, authorIDs []int64) (*postgres.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", ctx, title, authorIDs)
	ret0, _ := ret[0].(*postgres.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockRepositoryMockRecorder) CreateBook(ctx, title, authorIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockRepository)(nil).CreateBook), ctx, title, authorIDs)
}

// ListAuthors mocks base method.
func (m *MockRepository) ListAuthors(ctx context.Context) ([]postgres.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthors", ctx)
	ret0, _ := ret[0].([]postgres.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthors indicates an expected call of ListAuthors.
func (mr *MockRepositoryMockRecorder) ListAuthors(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthors", reflect.TypeOf((*MockRepository)(nil).ListAuthors), ctx)
}

// ListAuthorsByBookID mocks base method.
func (m *MockRepository) ListAuthorsByBookID(ctx context.Context, bookID int64) ([]postgres.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthorsByBookID", ctx, bookID)
	ret0, _ := ret[0].([]postgres.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAuthorsByBookID indicates an expected call of ListAuthorsByBookID.
func (mr *MockRepositoryMockRecorder) ListAuthorsByBookID(ctx, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthorsByBookID", reflect.TypeOf((*MockRepository)(nil).ListAuthorsByBookID), ctx, bookID)
}

// ListBooks mocks base method.
func (m *MockRepository) ListBooks(ctx context.Context) ([]postgres.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBooks", ctx)
	ret0, _ := ret[0].([]postgres.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBooks indicates an expected call of ListBooks.
func (mr *MockRepositoryMockRecorder) ListBooks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBooks", reflect.TypeOf((*MockRepository)(nil).ListBooks), ctx)
}

// ListBooksByAuthorID mocks base method.
func (m *MockRepository) ListBooksByAuthorID(ctx context.Context, authorID int64) ([]postgres.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBooksByAuthorID", ctx, authorID)
	ret0, _ := ret[0].([]postgres.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBooksByAuthorID indicates an expected call of ListBooksByAuthorID.
func (mr *MockRepositoryMockRecorder) ListBooksByAuthorID(ctx, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBooksByAuthorID", reflect.TypeOf((*MockRepository)(nil).ListBooksByAuthorID), ctx, authorID)
}

// UpdateAuthor mocks base method.
func (m *MockRepository) UpdateAuthor(ctx context.Context, arg postgres.UpdateAuthorParams) (postgres.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAuthor", ctx, arg)
	ret0, _ := ret[0].(postgres.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAuthor indicates an expected call of UpdateAuthor.
func (mr *MockRepositoryMockRecorder) UpdateAuthor(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAuthor", reflect.TypeOf((*MockRepository)(nil).UpdateAuthor), ctx, arg)
}

// UpdateBook mocks base method.
func (m *MockRepository) UpdateBook(ctx context.Context, bookArg postgres.UpdateBookParams, authorIDs []int64) (*postgres.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBook", ctx, bookArg, authorIDs)
	ret0, _ := ret[0].(*postgres.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBook indicates an expected call of UpdateBook.
func (mr *MockRepositoryMockRecorder) UpdateBook(ctx, bookArg, authorIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBook", reflect.TypeOf((*MockRepository)(nil).UpdateBook), ctx, bookArg, authorIDs)
}
