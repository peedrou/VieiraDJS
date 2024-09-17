package mocks

import (
	reflect "reflect"

	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
)

// MockSessionCreator is a mock of SessionCreator interface.
type MockSessionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSessionCreatorMockRecorder
}

// MockSessionCreatorMockRecorder is the mock recorder for MockSessionCreator.
type MockSessionCreatorMockRecorder struct {
	mock *MockSessionCreator
}

// NewMockSessionCreator creates a new mock instance.
func NewMockSessionCreator(ctrl *gomock.Controller) *MockSessionCreator {
	mock := &MockSessionCreator{ctrl: ctrl}
	mock.recorder = &MockSessionCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionCreator) EXPECT() *MockSessionCreatorMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockSessionCreator) CreateSession() (*gocql.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession")
	ret0, _ := ret[0].(*gocql.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionCreatorMockRecorder) CreateSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionCreator)(nil).CreateSession))
}

