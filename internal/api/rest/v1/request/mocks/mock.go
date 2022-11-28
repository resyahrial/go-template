// Code generated by MockGen. DO NOT EDIT.
// Source: adapter.go

// Package mock_request is a generated GoMock package.
package mock_request

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockValidator) Validate(data interface{}) map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", data)
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockValidatorMockRecorder) Validate(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockValidator)(nil).Validate), data)
}

// MockDecoder is a mock of Decoder interface.
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder.
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance.
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockDecoder) Decode(in, out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", in, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockDecoderMockRecorder) Decode(in, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode), in, out)
}

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// BindJSON mocks base method.
func (m *MockContext) BindJSON(obj any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindJSON", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// BindJSON indicates an expected call of BindJSON.
func (mr *MockContextMockRecorder) BindJSON(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindJSON", reflect.TypeOf((*MockContext)(nil).BindJSON), obj)
}