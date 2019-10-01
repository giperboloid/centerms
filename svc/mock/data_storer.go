// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kostiamol/centerms/svc (interfaces: DataStorer)

// Package mock_svc is a generated GoMock package.
package mock_svc

import (
	gomock "github.com/golang/mock/gomock"
	svc "github.com/kostiamol/centerms/svc"
	reflect "reflect"
)

// MockDataStorer is a mock of DataStorer interface
type MockDataStorer struct {
	ctrl     *gomock.Controller
	recorder *MockDataStorerMockRecorder
}

// MockDataStorerMockRecorder is the mock recorder for MockDataStorer
type MockDataStorerMockRecorder struct {
	mock *MockDataStorer
}

// NewMockDataStorer creates a new mock instance
func NewMockDataStorer(ctrl *gomock.Controller) *MockDataStorer {
	mock := &MockDataStorer{ctrl: ctrl}
	mock.recorder = &MockDataStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStorer) EXPECT() *MockDataStorerMockRecorder {
	return m.recorder
}

// GetDevData mocks base method
func (m *MockDataStorer) GetDevData(arg0 string) (*svc.DevData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevData", arg0)
	ret0, _ := ret[0].(*svc.DevData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevData indicates an expected call of GetDevData
func (mr *MockDataStorerMockRecorder) GetDevData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevData", reflect.TypeOf((*MockDataStorer)(nil).GetDevData), arg0)
}

// GetDevsData mocks base method
func (m *MockDataStorer) GetDevsData() ([]svc.DevData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevsData")
	ret0, _ := ret[0].([]svc.DevData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevsData indicates an expected call of GetDevsData
func (mr *MockDataStorerMockRecorder) GetDevsData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevsData", reflect.TypeOf((*MockDataStorer)(nil).GetDevsData))
}

// Publish mocks base method
func (m *MockDataStorer) Publish(arg0 interface{}, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish
func (mr *MockDataStorerMockRecorder) Publish(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockDataStorer)(nil).Publish), arg0, arg1)
}

// SaveDevData mocks base method
func (m *MockDataStorer) SaveDevData(arg0 *svc.DevData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDevData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDevData indicates an expected call of SaveDevData
func (mr *MockDataStorerMockRecorder) SaveDevData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDevData", reflect.TypeOf((*MockDataStorer)(nil).SaveDevData), arg0)
}
