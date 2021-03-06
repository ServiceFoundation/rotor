// Code generated by MockGen. DO NOT EDIT.
// Source: fromflags.go

package updater

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/turbinelabs/api"
	service "github.com/turbinelabs/api/service"
	poller "github.com/turbinelabs/rotor/xds/poller"
	reflect "reflect"
)

// MockFromFlags is a mock of FromFlags interface
type MockFromFlags struct {
	ctrl     *gomock.Controller
	recorder *MockFromFlagsMockRecorder
}

// MockFromFlagsMockRecorder is the mock recorder for MockFromFlags
type MockFromFlagsMockRecorder struct {
	mock *MockFromFlags
}

// NewMockFromFlags creates a new mock instance
func NewMockFromFlags(ctrl *gomock.Controller) *MockFromFlags {
	mock := &MockFromFlags{ctrl: ctrl}
	mock.recorder = &MockFromFlagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFromFlags) EXPECT() *MockFromFlagsMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockFromFlags) Validate() error {
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockFromFlagsMockRecorder) Validate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockFromFlags)(nil).Validate))
}

// Make mocks base method
func (m *MockFromFlags) Make(arg0 service.All, arg1 api.Zone) Updater {
	ret := m.ctrl.Call(m, "Make", arg0, arg1)
	ret0, _ := ret[0].(Updater)
	return ret0
}

// Make indicates an expected call of Make
func (mr *MockFromFlagsMockRecorder) Make(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Make", reflect.TypeOf((*MockFromFlags)(nil).Make), arg0, arg1)
}

// MakeStandalone mocks base method
func (m *MockFromFlags) MakeStandalone(port int) (func(poller.Consumer) Updater, poller.Registrar) {
	ret := m.ctrl.Call(m, "MakeStandalone", port)
	ret0, _ := ret[0].(func(poller.Consumer) Updater)
	ret1, _ := ret[1].(poller.Registrar)
	return ret0, ret1
}

// MakeStandalone indicates an expected call of MakeStandalone
func (mr *MockFromFlagsMockRecorder) MakeStandalone(port interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStandalone", reflect.TypeOf((*MockFromFlags)(nil).MakeStandalone), port)
}
