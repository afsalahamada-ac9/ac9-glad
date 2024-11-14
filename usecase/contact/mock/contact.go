// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_contact is a generated GoMock package.
package mock_contact

import (
	reflect "reflect"
	entity "sudhagar/glad/entity"

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

// GetByHandle mocks base method.
func (m *MockReader) GetByHandle(tenantID, accountID entity.ID, handle string) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHandle", tenantID, accountID, handle)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHandle indicates an expected call of GetByHandle.
func (mr *MockReaderMockRecorder) GetByHandle(tenantID, accountID, handle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHandle", reflect.TypeOf((*MockReader)(nil).GetByHandle), tenantID, accountID, handle)
}

// GetByID mocks base method.
func (m *MockReader) GetByID(tenantID, contactID entity.ID) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", tenantID, contactID)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockReaderMockRecorder) GetByID(tenantID, contactID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockReader)(nil).GetByID), tenantID, contactID)
}

// GetCount mocks base method.
func (m *MockReader) GetCount(tenantId entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockReaderMockRecorder) GetCount(tenantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockReader)(nil).GetCount), tenantId)
}

// GetMulti mocks base method.
func (m *MockReader) GetMulti(tenantID entity.ID, page, page_size int) ([]*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMulti", tenantID, page, page_size)
	ret0, _ := ret[0].([]*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockReaderMockRecorder) GetMulti(tenantID, page, page_size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockReader)(nil).GetMulti), tenantID, page, page_size)
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
func (m *MockWriter) Create(e *entity.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWriterMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWriter)(nil).Create), e)
}

// DeleteByAccount mocks base method.
func (m *MockWriter) DeleteByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByAccount indicates an expected call of DeleteByAccount.
func (mr *MockWriterMockRecorder) DeleteByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByAccount", reflect.TypeOf((*MockWriter)(nil).DeleteByAccount), tenantID, accountID)
}

// DeleteStaleByAccount mocks base method.
func (m *MockWriter) DeleteStaleByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStaleByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStaleByAccount indicates an expected call of DeleteStaleByAccount.
func (mr *MockWriterMockRecorder) DeleteStaleByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStaleByAccount", reflect.TypeOf((*MockWriter)(nil).DeleteStaleByAccount), tenantID, accountID)
}

// SetStaleByAccount mocks base method.
func (m *MockWriter) SetStaleByAccount(tenantID, accountID entity.ID, value bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStaleByAccount", tenantID, accountID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStaleByAccount indicates an expected call of SetStaleByAccount.
func (mr *MockWriterMockRecorder) SetStaleByAccount(tenantID, accountID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStaleByAccount", reflect.TypeOf((*MockWriter)(nil).SetStaleByAccount), tenantID, accountID, value)
}

// Update mocks base method.
func (m *MockWriter) Update(e *entity.Contact) error {
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
func (m *MockRepository) Create(e *entity.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), e)
}

// DeleteByAccount mocks base method.
func (m *MockRepository) DeleteByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByAccount indicates an expected call of DeleteByAccount.
func (mr *MockRepositoryMockRecorder) DeleteByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByAccount", reflect.TypeOf((*MockRepository)(nil).DeleteByAccount), tenantID, accountID)
}

// DeleteStaleByAccount mocks base method.
func (m *MockRepository) DeleteStaleByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStaleByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStaleByAccount indicates an expected call of DeleteStaleByAccount.
func (mr *MockRepositoryMockRecorder) DeleteStaleByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStaleByAccount", reflect.TypeOf((*MockRepository)(nil).DeleteStaleByAccount), tenantID, accountID)
}

// GetByHandle mocks base method.
func (m *MockRepository) GetByHandle(tenantID, accountID entity.ID, handle string) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHandle", tenantID, accountID, handle)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHandle indicates an expected call of GetByHandle.
func (mr *MockRepositoryMockRecorder) GetByHandle(tenantID, accountID, handle interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHandle", reflect.TypeOf((*MockRepository)(nil).GetByHandle), tenantID, accountID, handle)
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(tenantID, contactID entity.ID) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", tenantID, contactID)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(tenantID, contactID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), tenantID, contactID)
}

// GetCount mocks base method.
func (m *MockRepository) GetCount(tenantId entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockRepositoryMockRecorder) GetCount(tenantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockRepository)(nil).GetCount), tenantId)
}

// GetMulti mocks base method.
func (m *MockRepository) GetMulti(tenantID entity.ID, page, page_size int) ([]*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMulti", tenantID, page, page_size)
	ret0, _ := ret[0].([]*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockRepositoryMockRecorder) GetMulti(tenantID, page, page_size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockRepository)(nil).GetMulti), tenantID, page, page_size)
}

// SetStaleByAccount mocks base method.
func (m *MockRepository) SetStaleByAccount(tenantID, accountID entity.ID, value bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStaleByAccount", tenantID, accountID, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStaleByAccount indicates an expected call of SetStaleByAccount.
func (mr *MockRepositoryMockRecorder) SetStaleByAccount(tenantID, accountID, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStaleByAccount", reflect.TypeOf((*MockRepository)(nil).SetStaleByAccount), tenantID, accountID, value)
}

// Update mocks base method.
func (m *MockRepository) Update(e *entity.Contact) error {
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

// CreateOrUpdate mocks base method.
func (m *MockUseCase) CreateOrUpdate(tenantID, accountID entity.ID, handle, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", tenantID, accountID, handle, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockUseCaseMockRecorder) CreateOrUpdate(tenantID, accountID, handle, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockUseCase)(nil).CreateOrUpdate), tenantID, accountID, handle, name)
}

// DeleteByAccount mocks base method.
func (m *MockUseCase) DeleteByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByAccount indicates an expected call of DeleteByAccount.
func (mr *MockUseCaseMockRecorder) DeleteByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByAccount", reflect.TypeOf((*MockUseCase)(nil).DeleteByAccount), tenantID, accountID)
}

// DeleteStaleByAccount mocks base method.
func (m *MockUseCase) DeleteStaleByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStaleByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStaleByAccount indicates an expected call of DeleteStaleByAccount.
func (mr *MockUseCaseMockRecorder) DeleteStaleByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStaleByAccount", reflect.TypeOf((*MockUseCase)(nil).DeleteStaleByAccount), tenantID, accountID)
}

// Get mocks base method.
func (m *MockUseCase) Get(tenantID, contactID entity.ID) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", tenantID, contactID)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUseCaseMockRecorder) Get(tenantID, contactID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUseCase)(nil).Get), tenantID, contactID)
}

// GetCount mocks base method.
func (m *MockUseCase) GetCount(tenantId entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockUseCaseMockRecorder) GetCount(tenantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockUseCase)(nil).GetCount), tenantId)
}

// GetMulti mocks base method.
func (m *MockUseCase) GetMulti(tenantID entity.ID, page, page_size int) ([]*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMulti", tenantID, page, page_size)
	ret0, _ := ret[0].([]*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMulti indicates an expected call of GetMulti.
func (mr *MockUseCaseMockRecorder) GetMulti(tenantID, page, page_size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMulti", reflect.TypeOf((*MockUseCase)(nil).GetMulti), tenantID, page, page_size)
}

// ResetStaleByAccount mocks base method.
func (m *MockUseCase) ResetStaleByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetStaleByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetStaleByAccount indicates an expected call of ResetStaleByAccount.
func (mr *MockUseCaseMockRecorder) ResetStaleByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStaleByAccount", reflect.TypeOf((*MockUseCase)(nil).ResetStaleByAccount), tenantID, accountID)
}

// SetStaleByAccount mocks base method.
func (m *MockUseCase) SetStaleByAccount(tenantID, accountID entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStaleByAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStaleByAccount indicates an expected call of SetStaleByAccount.
func (mr *MockUseCaseMockRecorder) SetStaleByAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStaleByAccount", reflect.TypeOf((*MockUseCase)(nil).SetStaleByAccount), tenantID, accountID)
}

// Update mocks base method.
func (m *MockUseCase) Update(c *entity.Contact) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUseCaseMockRecorder) Update(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), c)
}
