// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/alexandre-slp/event-broker/app/api (interfaces: RouterInt)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	router "github.com/fasthttp/router"
	gomock "github.com/golang/mock/gomock"
)

// MockRouterInt is a mock of RouterInt interface.
type MockRouterInt struct {
	ctrl     *gomock.Controller
	recorder *MockRouterIntMockRecorder
}

// MockRouterIntMockRecorder is the mock recorder for MockRouterInt.
type MockRouterIntMockRecorder struct {
	mock *MockRouterInt
}

// NewMockRouterInt creates a new mock instance.
func NewMockRouterInt(ctrl *gomock.Controller) *MockRouterInt {
	mock := &MockRouterInt{ctrl: ctrl}
	mock.recorder = &MockRouterIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterInt) EXPECT() *MockRouterIntMockRecorder {
	return m.recorder
}

// InitPaths mocks base method.
func (m *MockRouterInt) InitPaths(arg0 *router.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitPaths", arg0)
}

// InitPaths indicates an expected call of InitPaths.
func (mr *MockRouterIntMockRecorder) InitPaths(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitPaths", reflect.TypeOf((*MockRouterInt)(nil).InitPaths), arg0)
}