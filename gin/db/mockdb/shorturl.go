// Code generated by MockGen. DO NOT EDIT.
// Source: shorturl/m/db/store/shorturl (interfaces: ShorturlQuery)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	sql "database/sql"
	reflect "reflect"
	shorturl "shorturl/m/db/store/shorturl"

	gomock "github.com/golang/mock/gomock"
)

// MockShorturlQuery is a mock of ShorturlQuery interface.
type MockShorturlQuery struct {
	ctrl     *gomock.Controller
	recorder *MockShorturlQueryMockRecorder
}

// MockShorturlQueryMockRecorder is the mock recorder for MockShorturlQuery.
type MockShorturlQueryMockRecorder struct {
	mock *MockShorturlQuery
}

// NewMockShorturlQuery creates a new mock instance.
func NewMockShorturlQuery(ctrl *gomock.Controller) *MockShorturlQuery {
	mock := &MockShorturlQuery{ctrl: ctrl}
	mock.recorder = &MockShorturlQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShorturlQuery) EXPECT() *MockShorturlQueryMockRecorder {
	return m.recorder
}

// CountMatchShorturl mocks base method.
func (m *MockShorturlQuery) CountMatchShorturl(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountMatchShorturl", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountMatchShorturl indicates an expected call of CountMatchShorturl.
func (mr *MockShorturlQueryMockRecorder) CountMatchShorturl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountMatchShorturl", reflect.TypeOf((*MockShorturlQuery)(nil).CountMatchShorturl), arg0, arg1)
}

// CreateShorturl mocks base method.
func (m *MockShorturlQuery) CreateShorturl(arg0 context.Context, arg1 shorturl.CreateShorturlParams) (shorturl.Shorturl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShorturl", arg0, arg1)
	ret0, _ := ret[0].(shorturl.Shorturl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShorturl indicates an expected call of CreateShorturl.
func (mr *MockShorturlQueryMockRecorder) CreateShorturl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShorturl", reflect.TypeOf((*MockShorturlQuery)(nil).CreateShorturl), arg0, arg1)
}

// DeleteShorturl mocks base method.
func (m *MockShorturlQuery) DeleteShorturl(arg0 context.Context, arg1 shorturl.DeleteShorturlParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShorturl", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteShorturl indicates an expected call of DeleteShorturl.
func (mr *MockShorturlQueryMockRecorder) DeleteShorturl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShorturl", reflect.TypeOf((*MockShorturlQuery)(nil).DeleteShorturl), arg0, arg1)
}

// GetMatchShorturl mocks base method.
func (m *MockShorturlQuery) GetMatchShorturl(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchShorturl", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMatchShorturl indicates an expected call of GetMatchShorturl.
func (mr *MockShorturlQueryMockRecorder) GetMatchShorturl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchShorturl", reflect.TypeOf((*MockShorturlQuery)(nil).GetMatchShorturl), arg0, arg1)
}

// ListUserShorturl mocks base method.
func (m *MockShorturlQuery) ListUserShorturl(arg0 context.Context, arg1 sql.NullInt32) ([]shorturl.ListUserShorturlRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserShorturl", arg0, arg1)
	ret0, _ := ret[0].([]shorturl.ListUserShorturlRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserShorturl indicates an expected call of ListUserShorturl.
func (mr *MockShorturlQueryMockRecorder) ListUserShorturl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserShorturl", reflect.TypeOf((*MockShorturlQuery)(nil).ListUserShorturl), arg0, arg1)
}

// UpdateExpired mocks base method.
func (m *MockShorturlQuery) UpdateExpired(arg0 context.Context, arg1 shorturl.UpdateExpiredParams) (shorturl.Shorturl, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpired", arg0, arg1)
	ret0, _ := ret[0].(shorturl.Shorturl)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateExpired indicates an expected call of UpdateExpired.
func (mr *MockShorturlQueryMockRecorder) UpdateExpired(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpired", reflect.TypeOf((*MockShorturlQuery)(nil).UpdateExpired), arg0, arg1)
}