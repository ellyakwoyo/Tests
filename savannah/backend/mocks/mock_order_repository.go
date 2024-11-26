// Code generated by MockGen. DO NOT EDIT.
// Source: backend/internal/repositories (interfaces: OrderRepositoryImpl)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "backend/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepositoryImpl is a mock of OrderRepositoryImpl interface.
type MockOrderRepositoryImpl struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryImplMockRecorder
}

// MockOrderRepositoryImplMockRecorder is the mock recorder for MockOrderRepositoryImpl.
type MockOrderRepositoryImplMockRecorder struct {
	mock *MockOrderRepositoryImpl
}

// NewMockOrderRepositoryImpl creates a new mock instance.
func NewMockOrderRepositoryImpl(ctrl *gomock.Controller) *MockOrderRepositoryImpl {
	mock := &MockOrderRepositoryImpl{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepositoryImpl) EXPECT() *MockOrderRepositoryImplMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrderRepositoryImpl) Create(arg0 *models.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOrderRepositoryImplMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockOrderRepositoryImpl) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrderRepositoryImplMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockOrderRepositoryImpl) GetAll() ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockOrderRepositoryImplMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockOrderRepositoryImpl) GetByID(arg0 int) (*models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockOrderRepositoryImplMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).GetByID), arg0)
}

// GetOrdersByUserID mocks base method.
func (m *MockOrderRepositoryImpl) GetOrdersByUserID(arg0 int) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByUserID", arg0)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByUserID indicates an expected call of GetOrdersByUserID.
func (mr *MockOrderRepositoryImplMockRecorder) GetOrdersByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByUserID", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).GetOrdersByUserID), arg0)
}

// Update mocks base method.
func (m *MockOrderRepositoryImpl) Update(arg0 *models.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOrderRepositoryImplMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOrderRepositoryImpl)(nil).Update), arg0)
}