// Code generated by MockGen. DO NOT EDIT.
// Source: client/client.go

// Package mock_client is a generated GoMock package.
package mock

import (
	model "bootcamp/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIClient is a mock of IClient interface.
type MockIClient struct {
	ctrl     *gomock.Controller
	recorder *MockIClientMockRecorder
}

// MockIClientMockRecorder is the mock recorder for MockIClient.
type MockIClientMockRecorder struct {
	mock *MockIClient
}

// NewMockIClient creates a new mock instance.
func NewMockIClient(ctrl *gomock.Controller) *MockIClient {
	mock := &MockIClient{ctrl: ctrl}
	mock.recorder = &MockIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClient) EXPECT() *MockIClientMockRecorder {
	return m.recorder
}

// GetQuotes mocks base method.
func (m *MockIClient) GetQuotes() (*model.ClientResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuotes")
	ret0, _ := ret[0].(*model.ClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuotes indicates an expected call of GetQuotes.
func (mr *MockIClientMockRecorder) GetQuotes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuotes", reflect.TypeOf((*MockIClient)(nil).GetQuotes))
}
