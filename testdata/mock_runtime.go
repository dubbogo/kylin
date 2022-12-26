// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/arana-db/arana/pkg/proto (interfaces: VConn,Plan,Optimizer,DB)

// Package testdata is a generated GoMock package.
package testdata

import (
	context "context"
	reflect "reflect"
	time "time"
)

import (
	gomock "github.com/golang/mock/gomock"
)

import (
	proto "github.com/arana-db/arana/pkg/proto"
)

// MockVConn is a mock of VConn interface.
type MockVConn struct {
	ctrl     *gomock.Controller
	recorder *MockVConnMockRecorder
}

// MockVConnMockRecorder is the mock recorder for MockVConn.
type MockVConnMockRecorder struct {
	mock *MockVConn
}

// NewMockVConn creates a new mock instance.
func NewMockVConn(ctrl *gomock.Controller) *MockVConn {
	mock := &MockVConn{ctrl: ctrl}
	mock.recorder = &MockVConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVConn) EXPECT() *MockVConnMockRecorder {
	return m.recorder
}

// Exec mocks base method.
func (m *MockVConn) Exec(arg0 context.Context, arg1, arg2 string, arg3 ...proto.Value) (proto.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockVConnMockRecorder) Exec(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockVConn)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockVConn) Query(arg0 context.Context, arg1, arg2 string, arg3 ...proto.Value) (proto.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockVConnMockRecorder) Query(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockVConn)(nil).Query), varargs...)
}

// MockPlan is a mock of Plan interface.
type MockPlan struct {
	ctrl     *gomock.Controller
	recorder *MockPlanMockRecorder
}

// MockPlanMockRecorder is the mock recorder for MockPlan.
type MockPlanMockRecorder struct {
	mock *MockPlan
}

// NewMockPlan creates a new mock instance.
func NewMockPlan(ctrl *gomock.Controller) *MockPlan {
	mock := &MockPlan{ctrl: ctrl}
	mock.recorder = &MockPlanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlan) EXPECT() *MockPlanMockRecorder {
	return m.recorder
}

// ExecIn mocks base method.
func (m *MockPlan) ExecIn(arg0 context.Context, arg1 proto.VConn) (proto.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecIn", arg0, arg1)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecIn indicates an expected call of ExecIn.
func (mr *MockPlanMockRecorder) ExecIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecIn", reflect.TypeOf((*MockPlan)(nil).ExecIn), arg0, arg1)
}

// Type mocks base method.
func (m *MockPlan) Type() proto.PlanType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(proto.PlanType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockPlanMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockPlan)(nil).Type))
}

// MockOptimizer is a mock of Optimizer interface.
type MockOptimizer struct {
	ctrl     *gomock.Controller
	recorder *MockOptimizerMockRecorder
}

// MockOptimizerMockRecorder is the mock recorder for MockOptimizer.
type MockOptimizerMockRecorder struct {
	mock *MockOptimizer
}

// NewMockOptimizer creates a new mock instance.
func NewMockOptimizer(ctrl *gomock.Controller) *MockOptimizer {
	mock := &MockOptimizer{ctrl: ctrl}
	mock.recorder = &MockOptimizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOptimizer) EXPECT() *MockOptimizerMockRecorder {
	return m.recorder
}

// Optimize mocks base method.
func (m *MockOptimizer) Optimize(arg0 context.Context) (proto.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Optimize", arg0)
	ret0, _ := ret[0].(proto.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Optimize indicates an expected call of Optimize.
func (mr *MockOptimizerMockRecorder) Optimize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Optimize", reflect.TypeOf((*MockOptimizer)(nil).Optimize), arg0)
}

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockDB) Call(arg0 context.Context, arg1 string, arg2 ...proto.Value) (proto.Result, uint16, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(proto.Result)
	ret1, _ := ret[1].(uint16)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Call indicates an expected call of Call.
func (mr *MockDBMockRecorder) Call(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockDB)(nil).Call), varargs...)
}

// CallFieldList mocks base method.
func (m *MockDB) CallFieldList(arg0 context.Context, arg1, arg2 string) ([]proto.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallFieldList", arg0, arg1, arg2)
	ret0, _ := ret[0].([]proto.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallFieldList indicates an expected call of CallFieldList.
func (mr *MockDBMockRecorder) CallFieldList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallFieldList", reflect.TypeOf((*MockDB)(nil).CallFieldList), arg0, arg1, arg2)
}

// Capacity mocks base method.
func (m *MockDB) Capacity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capacity")
	ret0, _ := ret[0].(int)
	return ret0
}

// Capacity indicates an expected call of Capacity.
func (mr *MockDBMockRecorder) Capacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capacity", reflect.TypeOf((*MockDB)(nil).Capacity))
}

// Close mocks base method.
func (m *MockDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}

// ID mocks base method.
func (m *MockDB) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockDBMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockDB)(nil).ID))
}

// IdleTimeout mocks base method.
func (m *MockDB) IdleTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IdleTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// IdleTimeout indicates an expected call of IdleTimeout.
func (mr *MockDBMockRecorder) IdleTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IdleTimeout", reflect.TypeOf((*MockDB)(nil).IdleTimeout))
}

// MaxCapacity mocks base method.
func (m *MockDB) MaxCapacity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxCapacity")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxCapacity indicates an expected call of MaxCapacity.
func (mr *MockDBMockRecorder) MaxCapacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxCapacity", reflect.TypeOf((*MockDB)(nil).MaxCapacity))
}

// SetCapacity mocks base method.
func (m *MockDB) SetCapacity(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCapacity", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCapacity indicates an expected call of SetCapacity.
func (mr *MockDBMockRecorder) SetCapacity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCapacity", reflect.TypeOf((*MockDB)(nil).SetCapacity), arg0)
}

// SetIdleTimeout mocks base method.
func (m *MockDB) SetIdleTimeout(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIdleTimeout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIdleTimeout indicates an expected call of SetIdleTimeout.
func (mr *MockDBMockRecorder) SetIdleTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIdleTimeout", reflect.TypeOf((*MockDB)(nil).SetIdleTimeout), arg0)
}

// SetMaxCapacity mocks base method.
func (m *MockDB) SetMaxCapacity(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMaxCapacity", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMaxCapacity indicates an expected call of SetMaxCapacity.
func (mr *MockDBMockRecorder) SetMaxCapacity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxCapacity", reflect.TypeOf((*MockDB)(nil).SetMaxCapacity), arg0)
}

// SetWeight mocks base method.
func (m *MockDB) SetWeight(arg0 proto.Weight) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWeight", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWeight indicates an expected call of SetWeight.
func (mr *MockDBMockRecorder) SetWeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWeight", reflect.TypeOf((*MockDB)(nil).SetWeight), arg0)
}

// Variable mocks base method.
func (m *MockDB) Variable(arg0 context.Context, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Variable", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Variable indicates an expected call of Variable.
func (mr *MockDBMockRecorder) Variable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Variable", reflect.TypeOf((*MockDB)(nil).Variable), arg0, arg1)
}

// Weight mocks base method.
func (m *MockDB) Weight() proto.Weight {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Weight")
	ret0, _ := ret[0].(proto.Weight)
	return ret0
}

// Weight indicates an expected call of Weight.
func (mr *MockDBMockRecorder) Weight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Weight", reflect.TypeOf((*MockDB)(nil).Weight))
}
