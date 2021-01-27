// Code generated by MockGen. DO NOT EDIT.
// Source: rule.go

// Package auth0 is a generated GoMock package.
package auth0

import (
	gomock "github.com/golang/mock/gomock"
	management "gopkg.in/auth0.v5/management"
	reflect "reflect"
)

// MockRuleAPI is a mock of RuleAPI interface
type MockRuleAPI struct {
	ctrl     *gomock.Controller
	recorder *MockRuleAPIMockRecorder
}

// MockRuleAPIMockRecorder is the mock recorder for MockRuleAPI
type MockRuleAPIMockRecorder struct {
	mock *MockRuleAPI
}

// NewMockRuleAPI creates a new mock instance
func NewMockRuleAPI(ctrl *gomock.Controller) *MockRuleAPI {
	mock := &MockRuleAPI{ctrl: ctrl}
	mock.recorder = &MockRuleAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRuleAPI) EXPECT() *MockRuleAPIMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRuleAPI) Create(r *management.Rule, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{r}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRuleAPIMockRecorder) Create(r interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{r}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRuleAPI)(nil).Create), varargs...)
}

// Read mocks base method
func (m *MockRuleAPI) Read(id string, opts ...management.RequestOption) (*management.Rule, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Read", varargs...)
	ret0, _ := ret[0].(*management.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockRuleAPIMockRecorder) Read(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockRuleAPI)(nil).Read), varargs...)
}

// Update mocks base method
func (m *MockRuleAPI) Update(id string, r *management.Rule, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{id, r}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRuleAPIMockRecorder) Update(id, r interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id, r}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRuleAPI)(nil).Update), varargs...)
}

// Delete mocks base method
func (m *MockRuleAPI) Delete(id string, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockRuleAPIMockRecorder) Delete(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRuleAPI)(nil).Delete), varargs...)
}

// List mocks base method
func (m *MockRuleAPI) List(opts ...management.RequestOption) (*management.RuleList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*management.RuleList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRuleAPIMockRecorder) List(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRuleAPI)(nil).List), opts...)
}