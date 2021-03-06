// Code generated by MockGen. DO NOT EDIT.
// Source: getir-assignment/internal/record (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	record "getir-assignment/internal/record"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetRecordsByDateAndCount mocks base method.
func (m *MockRepository) GetRecordsByDateAndCount(arg0, arg1 time.Time, arg2, arg3 int) ([]*record.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordsByDateAndCount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*record.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordsByDateAndCount indicates an expected call of GetRecordsByDateAndCount.
func (mr *MockRepositoryMockRecorder) GetRecordsByDateAndCount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsByDateAndCount", reflect.TypeOf((*MockRepository)(nil).GetRecordsByDateAndCount), arg0, arg1, arg2, arg3)
}
