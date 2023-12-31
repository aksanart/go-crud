// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/main.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "aksan-go-crud/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceIntf is a mock of ServiceIntf interface.
type MockServiceIntf struct {
	ctrl     *gomock.Controller
	recorder *MockServiceIntfMockRecorder
}

// MockServiceIntfMockRecorder is the mock recorder for MockServiceIntf.
type MockServiceIntfMockRecorder struct {
	mock *MockServiceIntf
}

// NewMockServiceIntf creates a new mock instance.
func NewMockServiceIntf(ctrl *gomock.Controller) *MockServiceIntf {
	mock := &MockServiceIntf{ctrl: ctrl}
	mock.recorder = &MockServiceIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceIntf) EXPECT() *MockServiceIntfMockRecorder {
	return m.recorder
}

// GetService mocks base method.
func (m *MockServiceIntf) GetService() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockServiceIntfMockRecorder) GetService() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockServiceIntf)(nil).GetService))
}

// SaveService mocks base method.
func (m *MockServiceIntf) SaveService(req *entity.SaveRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveService", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveService indicates an expected call of SaveService.
func (mr *MockServiceIntfMockRecorder) SaveService(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveService", reflect.TypeOf((*MockServiceIntf)(nil).SaveService), req)
}
