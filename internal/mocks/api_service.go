// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/api/service (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/base-org/pessimism/internal/api/models"
	core "github.com/base-org/pessimism/internal/core"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CheckETHRPCHealth mocks base method.
func (m *MockService) CheckETHRPCHealth(arg0 core.Network) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckETHRPCHealth", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckETHRPCHealth indicates an expected call of CheckETHRPCHealth.
func (mr *MockServiceMockRecorder) CheckETHRPCHealth(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckETHRPCHealth", reflect.TypeOf((*MockService)(nil).CheckETHRPCHealth), arg0)
}

// CheckHealth mocks base method.
func (m *MockService) CheckHealth() *models.HealthCheck {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth")
	ret0, _ := ret[0].(*models.HealthCheck)
	return ret0
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockServiceMockRecorder) CheckHealth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockService)(nil).CheckHealth))
}

// ProcessInvariantRequest mocks base method.
func (m *MockService) ProcessInvariantRequest(arg0 *models.InvRequestBody) (core.SUUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessInvariantRequest", arg0)
	ret0, _ := ret[0].(core.SUUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessInvariantRequest indicates an expected call of ProcessInvariantRequest.
func (mr *MockServiceMockRecorder) ProcessInvariantRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessInvariantRequest", reflect.TypeOf((*MockService)(nil).ProcessInvariantRequest), arg0)
}

// RunInvariantSession mocks base method.
func (m *MockService) RunInvariantSession(arg0 *models.InvRequestParams) (core.SUUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunInvariantSession", arg0)
	ret0, _ := ret[0].(core.SUUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunInvariantSession indicates an expected call of RunInvariantSession.
func (mr *MockServiceMockRecorder) RunInvariantSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInvariantSession", reflect.TypeOf((*MockService)(nil).RunInvariantSession), arg0)
}
