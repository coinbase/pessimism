// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/client (interfaces: EthClientInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEthClientInterface is a mock of EthClientInterface interface.
type MockEthClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEthClientInterfaceMockRecorder
}

// MockEthClientInterfaceMockRecorder is the mock recorder for MockEthClientInterface.
type MockEthClientInterfaceMockRecorder struct {
	mock *MockEthClientInterface
}

// NewMockEthClientInterface creates a new mock instance.
func NewMockEthClientInterface(ctrl *gomock.Controller) *MockEthClientInterface {
	mock := &MockEthClientInterface{ctrl: ctrl}
	mock.recorder = &MockEthClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthClientInterface) EXPECT() *MockEthClientInterfaceMockRecorder {
	return m.recorder
}

// BlockByNumber mocks base method.
func (m *MockEthClientInterface) BlockByNumber(arg0 context.Context, arg1 *big.Int) (*types.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockEthClientInterfaceMockRecorder) BlockByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockEthClientInterface)(nil).BlockByNumber), arg0, arg1)
}

// DialContext mocks base method.
func (m *MockEthClientInterface) DialContext(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DialContext indicates an expected call of DialContext.
func (mr *MockEthClientInterfaceMockRecorder) DialContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialContext", reflect.TypeOf((*MockEthClientInterface)(nil).DialContext), arg0, arg1)
}

// HeaderByNumber mocks base method.
func (m *MockEthClientInterface) HeaderByNumber(arg0 context.Context, arg1 *big.Int) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockEthClientInterfaceMockRecorder) HeaderByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockEthClientInterface)(nil).HeaderByNumber), arg0, arg1)
}
