// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/device/interface.go

// Package mock_device is a generated GoMock package.
package mock_device

import (
	id "ac9/glad/pkg/id"
	entity "ac9/glad/services/pushd/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeviceReader is a mock of DeviceReader interface.
type MockDeviceReader struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceReaderMockRecorder
}

// MockDeviceReaderMockRecorder is the mock recorder for MockDeviceReader.
type MockDeviceReaderMockRecorder struct {
	mock *MockDeviceReader
}

// NewMockDeviceReader creates a new mock instance.
func NewMockDeviceReader(ctrl *gomock.Controller) *MockDeviceReader {
	mock := &MockDeviceReader{ctrl: ctrl}
	mock.recorder = &MockDeviceReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceReader) EXPECT() *MockDeviceReaderMockRecorder {
	return m.recorder
}

// GetByAccount mocks base method.
func (m *MockDeviceReader) GetByAccount(tenantID, accountID id.ID) ([]*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccount", tenantID, accountID)
	ret0, _ := ret[0].([]*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAccount indicates an expected call of GetByAccount.
func (mr *MockDeviceReaderMockRecorder) GetByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccount", reflect.TypeOf((*MockDeviceReader)(nil).GetByAccount), tenantID, accountID)
}

// GetCount mocks base method.
func (m *MockDeviceReader) GetCount(tenantID id.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockDeviceReaderMockRecorder) GetCount(tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockDeviceReader)(nil).GetCount), tenantID)
}

// MockDeviceWriter is a mock of DeviceWriter interface.
type MockDeviceWriter struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceWriterMockRecorder
}

// MockDeviceWriterMockRecorder is the mock recorder for MockDeviceWriter.
type MockDeviceWriterMockRecorder struct {
	mock *MockDeviceWriter
}

// NewMockDeviceWriter creates a new mock instance.
func NewMockDeviceWriter(ctrl *gomock.Controller) *MockDeviceWriter {
	mock := &MockDeviceWriter{ctrl: ctrl}
	mock.recorder = &MockDeviceWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceWriter) EXPECT() *MockDeviceWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDeviceWriter) Create(e *entity.Device) (id.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDeviceWriterMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceWriter)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockDeviceWriter) Delete(id id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeviceWriterMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceWriter)(nil).Delete), id)
}

// MockDeviceRepository is a mock of DeviceRepository interface.
type MockDeviceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceRepositoryMockRecorder
}

// MockDeviceRepositoryMockRecorder is the mock recorder for MockDeviceRepository.
type MockDeviceRepositoryMockRecorder struct {
	mock *MockDeviceRepository
}

// NewMockDeviceRepository creates a new mock instance.
func NewMockDeviceRepository(ctrl *gomock.Controller) *MockDeviceRepository {
	mock := &MockDeviceRepository{ctrl: ctrl}
	mock.recorder = &MockDeviceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeviceRepository) EXPECT() *MockDeviceRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDeviceRepository) Create(e *entity.Device) (id.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDeviceRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDeviceRepository)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockDeviceRepository) Delete(id id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeviceRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeviceRepository)(nil).Delete), id)
}

// GetByAccount mocks base method.
func (m *MockDeviceRepository) GetByAccount(tenantID, accountID id.ID) ([]*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccount", tenantID, accountID)
	ret0, _ := ret[0].([]*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAccount indicates an expected call of GetByAccount.
func (mr *MockDeviceRepositoryMockRecorder) GetByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccount", reflect.TypeOf((*MockDeviceRepository)(nil).GetByAccount), tenantID, accountID)
}

// GetCount mocks base method.
func (m *MockDeviceRepository) GetCount(tenantID id.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockDeviceRepositoryMockRecorder) GetCount(tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockDeviceRepository)(nil).GetCount), tenantID)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateDevice mocks base method.
func (m *MockUseCase) CreateDevice(device entity.Device) (id.ID, []id.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDevice", device)
	ret0, _ := ret[0].(id.ID)
	ret1, _ := ret[1].([]id.ID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateDevice indicates an expected call of CreateDevice.
func (mr *MockUseCaseMockRecorder) CreateDevice(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDevice", reflect.TypeOf((*MockUseCase)(nil).CreateDevice), device)
}

// DeleteDevice mocks base method.
func (m *MockUseCase) DeleteDevice(id id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockUseCaseMockRecorder) DeleteDevice(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockUseCase)(nil).DeleteDevice), id)
}

// GetCount mocks base method.
func (m *MockUseCase) GetCount(id id.ID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", id)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCount indicates an expected call of GetCount.
func (mr *MockUseCaseMockRecorder) GetCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockUseCase)(nil).GetCount), id)
}

// GetDeviceByAccount mocks base method.
func (m *MockUseCase) GetDeviceByAccount(tenantID, accountID id.ID) ([]*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByAccount", tenantID, accountID)
	ret0, _ := ret[0].([]*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByAccount indicates an expected call of GetDeviceByAccount.
func (mr *MockUseCaseMockRecorder) GetDeviceByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByAccount", reflect.TypeOf((*MockUseCase)(nil).GetDeviceByAccount), tenantID, accountID)
}