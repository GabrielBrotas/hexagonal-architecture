// Code generated by MockGen. DO NOT EDIT.
// Source: application/product.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	application "github.com/GabrielBrotas/hexagonal-architeture/application"
	gomock "github.com/golang/mock/gomock"
)

// MockIProduct is a mock of IProduct interface.
type MockIProduct struct {
	ctrl     *gomock.Controller
	recorder *MockIProductMockRecorder
}

// MockIProductMockRecorder is the mock recorder for MockIProduct.
type MockIProductMockRecorder struct {
	mock *MockIProduct
}

// NewMockIProduct creates a new mock instance.
func NewMockIProduct(ctrl *gomock.Controller) *MockIProduct {
	mock := &MockIProduct{ctrl: ctrl}
	mock.recorder = &MockIProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProduct) EXPECT() *MockIProductMockRecorder {
	return m.recorder
}

// Disable mocks base method.
func (m *MockIProduct) Disable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockIProductMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIProduct)(nil).Disable))
}

// Enable mocks base method.
func (m *MockIProduct) Enable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockIProductMockRecorder) Enable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockIProduct)(nil).Enable))
}

// GetID mocks base method.
func (m *MockIProduct) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockIProductMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockIProduct)(nil).GetID))
}

// GetName mocks base method.
func (m *MockIProduct) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockIProductMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIProduct)(nil).GetName))
}

// GetPrice mocks base method.
func (m *MockIProduct) GetPrice() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockIProductMockRecorder) GetPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockIProduct)(nil).GetPrice))
}

// GetStatus mocks base method.
func (m *MockIProduct) GetStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockIProductMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockIProduct)(nil).GetStatus))
}

// IsValid mocks base method.
func (m *MockIProduct) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockIProductMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockIProduct)(nil).IsValid))
}

// MockIProductService is a mock of IProductService interface.
type MockIProductService struct {
	ctrl     *gomock.Controller
	recorder *MockIProductServiceMockRecorder
}

// MockIProductServiceMockRecorder is the mock recorder for MockIProductService.
type MockIProductServiceMockRecorder struct {
	mock *MockIProductService
}

// NewMockIProductService creates a new mock instance.
func NewMockIProductService(ctrl *gomock.Controller) *MockIProductService {
	mock := &MockIProductService{ctrl: ctrl}
	mock.recorder = &MockIProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductService) EXPECT() *MockIProductServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProductService) Create(name string, price float64) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIProductServiceMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProductService)(nil).Create), name, price)
}

// Disable mocks base method.
func (m *MockIProductService) Disable(product application.IProduct) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", product)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disable indicates an expected call of Disable.
func (mr *MockIProductServiceMockRecorder) Disable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIProductService)(nil).Disable), product)
}

// Enable mocks base method.
func (m *MockIProductService) Enable(product application.IProduct) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", product)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enable indicates an expected call of Enable.
func (mr *MockIProductServiceMockRecorder) Enable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockIProductService)(nil).Enable), product)
}

// GetByID mocks base method.
func (m *MockIProductService) GetByID(id string) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIProductServiceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIProductService)(nil).GetByID), id)
}

// MockIProductReader is a mock of IProductReader interface.
type MockIProductReader struct {
	ctrl     *gomock.Controller
	recorder *MockIProductReaderMockRecorder
}

// MockIProductReaderMockRecorder is the mock recorder for MockIProductReader.
type MockIProductReaderMockRecorder struct {
	mock *MockIProductReader
}

// NewMockIProductReader creates a new mock instance.
func NewMockIProductReader(ctrl *gomock.Controller) *MockIProductReader {
	mock := &MockIProductReader{ctrl: ctrl}
	mock.recorder = &MockIProductReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductReader) EXPECT() *MockIProductReaderMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockIProductReader) GetById(id string) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIProductReaderMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIProductReader)(nil).GetById), id)
}

// MockIProductWriter is a mock of IProductWriter interface.
type MockIProductWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIProductWriterMockRecorder
}

// MockIProductWriterMockRecorder is the mock recorder for MockIProductWriter.
type MockIProductWriterMockRecorder struct {
	mock *MockIProductWriter
}

// NewMockIProductWriter creates a new mock instance.
func NewMockIProductWriter(ctrl *gomock.Controller) *MockIProductWriter {
	mock := &MockIProductWriter{ctrl: ctrl}
	mock.recorder = &MockIProductWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductWriter) EXPECT() *MockIProductWriterMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockIProductWriter) Save(product application.IProduct) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIProductWriterMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIProductWriter)(nil).Save), product)
}

// MockIProductPersistance is a mock of IProductPersistance interface.
type MockIProductPersistance struct {
	ctrl     *gomock.Controller
	recorder *MockIProductPersistanceMockRecorder
}

// MockIProductPersistanceMockRecorder is the mock recorder for MockIProductPersistance.
type MockIProductPersistanceMockRecorder struct {
	mock *MockIProductPersistance
}

// NewMockIProductPersistance creates a new mock instance.
func NewMockIProductPersistance(ctrl *gomock.Controller) *MockIProductPersistance {
	mock := &MockIProductPersistance{ctrl: ctrl}
	mock.recorder = &MockIProductPersistanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductPersistance) EXPECT() *MockIProductPersistanceMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockIProductPersistance) GetById(id string) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIProductPersistanceMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIProductPersistance)(nil).GetById), id)
}

// Save mocks base method.
func (m *MockIProductPersistance) Save(product application.IProduct) (application.IProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.IProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIProductPersistanceMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIProductPersistance)(nil).Save), product)
}
