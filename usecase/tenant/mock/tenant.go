// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/tenant/interface.go

// Package mock_tenant is a generated GoMock package.
package mock_tenant

import (
	reflect "reflect"
	entity "ac9/glad/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockReader) Get(id entity.ID) (*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), id)
}

// GetByName mocks base method.
func (m *MockReader) GetByName(username string) (*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", username)
	ret0, _ := ret[0].(*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockReaderMockRecorder) GetByName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockReader)(nil).GetByName), username)
}

// GetCount mocks base method.
func (m *MockReader) GetCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockReaderMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockReader)(nil).GetCount))
}

// List mocks base method.
func (m *MockReader) List(page, limit int) ([]*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", page, limit)
	ret0, _ := ret[0].([]*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReaderMockRecorder) List(page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List), page, limit)
}

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter.
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWriter) Create(e *entity.Tenant) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWriterMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriter)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockWriter) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWriterMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), id)
}

// Update mocks base method.
func (m *MockWriter) Update(e *entity.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWriterMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWriter)(nil).Update), e)
}

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

// Create mocks base method.
func (m *MockRepository) Create(e *entity.Tenant) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// Delete mocks base method.
func (m *MockRepository) Delete(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockRepository) Get(id entity.ID) (*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// GetByName mocks base method.
func (m *MockRepository) GetByName(username string) (*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", username)
	ret0, _ := ret[0].(*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockRepositoryMockRecorder) GetByName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRepository)(nil).GetByName), username)
}

// GetCount mocks base method.
func (m *MockRepository) GetCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockRepositoryMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockRepository)(nil).GetCount))
}

// List mocks base method.
func (m *MockRepository) List(page, limit int) ([]*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", page, limit)
	ret0, _ := ret[0].([]*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), page, limit)
}

// Update mocks base method.
func (m *MockRepository) Update(e *entity.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), e)
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

// CreateTenant mocks base method.
func (m *MockUseCase) CreateTenant(username, country string) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTenant", username, country)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTenant indicates an expected call of CreateTenant.
func (mr *MockUseCaseMockRecorder) CreateTenant(username, country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTenant", reflect.TypeOf((*MockUseCase)(nil).CreateTenant), username, country)
}

// DeleteTenant mocks base method.
func (m *MockUseCase) DeleteTenant(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTenant", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTenant indicates an expected call of DeleteTenant.
func (mr *MockUseCaseMockRecorder) DeleteTenant(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTenant", reflect.TypeOf((*MockUseCase)(nil).DeleteTenant), id)
}

// GetCount mocks base method.
func (m *MockUseCase) GetCount() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCount indicates an expected call of GetCount.
func (mr *MockUseCaseMockRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockUseCase)(nil).GetCount))
}

// GetTenant mocks base method.
func (m *MockUseCase) GetTenant(id entity.ID) (*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenant", id)
	ret0, _ := ret[0].(*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenant indicates an expected call of GetTenant.
func (mr *MockUseCaseMockRecorder) GetTenant(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenant", reflect.TypeOf((*MockUseCase)(nil).GetTenant), id)
}

// ListTenants mocks base method.
func (m *MockUseCase) ListTenants(page, limit int) ([]*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTenants", page, limit)
	ret0, _ := ret[0].([]*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTenants indicates an expected call of ListTenants.
func (mr *MockUseCaseMockRecorder) ListTenants(page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTenants", reflect.TypeOf((*MockUseCase)(nil).ListTenants), page, limit)
}

// Login mocks base method.
func (m *MockUseCase) Login(username, password string) (*entity.Tenant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password)
	ret0, _ := ret[0].(*entity.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUseCaseMockRecorder) Login(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCase)(nil).Login), username, password)
}

// UpdateTenant mocks base method.
func (m *MockUseCase) UpdateTenant(e *entity.Tenant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTenant", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTenant indicates an expected call of UpdateTenant.
func (mr *MockUseCaseMockRecorder) UpdateTenant(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTenant", reflect.TypeOf((*MockUseCase)(nil).UpdateTenant), e)
}
