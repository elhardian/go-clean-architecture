// Package httpClient is a generated GoMock package.
package httpClient

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttp is a mock of Http interface.
type MockHttp struct {
	ctrl     *gomock.Controller
	recorder *MockHttpMockRecorder
}

// MockHttpMockRecorder is the mock recorder for MockHttp.
type MockHttpMockRecorder struct {
	mock *MockHttp
}

// NewMockHttp creates a new mock instance.
func NewMockHttp(ctrl *gomock.Controller) *MockHttp {
	mock := &MockHttp{ctrl: ctrl}
	mock.recorder = &MockHttpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttp) EXPECT() *MockHttpMockRecorder {
	return m.recorder
}

// CallURL mocks base method.
func (m *MockHttp) CallURL(method, url string, header map[string]string, rawData []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallURL", method, url, header, rawData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallURL indicates an expected call of CallURL.
func (mr *MockHttpMockRecorder) CallURL(method, url, header, rawData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallURL", reflect.TypeOf((*MockHttp)(nil).CallURL), method, url, header, rawData)
}

// Connect mocks base method.
func (m *MockHttp) Connect() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Connect")
}

// Connect indicates an expected call of Connect.
func (mr *MockHttpMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockHttp)(nil).Connect))
}
