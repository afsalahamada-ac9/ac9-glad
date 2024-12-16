// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/account/interface.go

// Package mock_account is a generated GoMock package.
package mock_account

import (
	entity "ac9/glad/entity"
	id "ac9/glad/pkg/id"
	reflect "reflect"

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
func (m *MockReader) Get(accountID id.ID) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", accountID)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReaderMockRecorder) Get(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), accountID)
}

// GetByEmail mocks base method.
func (m *MockReader) GetByEmail(tenantID id.ID, email string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", tenantID, email)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockReaderMockRecorder) GetByEmail(tenantID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockReader)(nil).GetByEmail), tenantID, email)
}

// GetByName mocks base method.
func (m *MockReader) GetByName(tenantID id.ID, username string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", tenantID, username)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockReaderMockRecorder) GetByName(tenantID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockReader)(nil).GetByName), tenantID, username)
}

// GetCount mocks base method.
func (m *MockReader) GetCount(tenantID id.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockReaderMockRecorder) GetCount(tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockReader)(nil).GetCount), tenantID)
}

// List mocks base method.
func (m *MockReader) List(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tenantID, page, limit, at)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReaderMockRecorder) List(tenantID, page, limit, at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List), tenantID, page, limit, at)
}

// Search mocks base method.
func (m *MockReader) Search(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", tenantID, query, page, limit, at)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockReaderMockRecorder) Search(tenantID, query, page, limit, at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockReader)(nil).Search), tenantID, query, page, limit, at)
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
func (m *MockWriter) Create(e *entity.Account) error {
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

// Delete mocks base method.
func (m *MockWriter) Delete(accountID id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWriterMockRecorder) Delete(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWriter)(nil).Delete), accountID)
}

// DeleteByName mocks base method.
func (m *MockWriter) DeleteByName(tenantID id.ID, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", tenantID, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockWriterMockRecorder) DeleteByName(tenantID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockWriter)(nil).DeleteByName), tenantID, username)
}

// Update mocks base method.
func (m *MockWriter) Update(e *entity.Account) error {
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

// Upsert mocks base method.
func (m *MockWriter) Upsert(e *entity.Account) (id.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", e)
	ret0, _ := ret[0].(id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockWriterMockRecorder) Upsert(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockWriter)(nil).Upsert), e)
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
func (m *MockRepository) Create(e *entity.Account) error {
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

// Delete mocks base method.
func (m *MockRepository) Delete(accountID id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), accountID)
}

// DeleteByName mocks base method.
func (m *MockRepository) DeleteByName(tenantID id.ID, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", tenantID, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockRepositoryMockRecorder) DeleteByName(tenantID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockRepository)(nil).DeleteByName), tenantID, username)
}

// Get mocks base method.
func (m *MockRepository) Get(accountID id.ID) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", accountID)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), accountID)
}

// GetByEmail mocks base method.
func (m *MockRepository) GetByEmail(tenantID id.ID, email string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", tenantID, email)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockRepositoryMockRecorder) GetByEmail(tenantID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockRepository)(nil).GetByEmail), tenantID, email)
}

// GetByName mocks base method.
func (m *MockRepository) GetByName(tenantID id.ID, username string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", tenantID, username)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockRepositoryMockRecorder) GetByName(tenantID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRepository)(nil).GetByName), tenantID, username)
}

// GetCount mocks base method.
func (m *MockRepository) GetCount(tenantID id.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockRepositoryMockRecorder) GetCount(tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockRepository)(nil).GetCount), tenantID)
}

// List mocks base method.
func (m *MockRepository) List(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tenantID, page, limit, at)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(tenantID, page, limit, at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), tenantID, page, limit, at)
}

// Search mocks base method.
func (m *MockRepository) Search(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", tenantID, query, page, limit, at)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRepositoryMockRecorder) Search(tenantID, query, page, limit, at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), tenantID, query, page, limit, at)
}

// Update mocks base method.
func (m *MockRepository) Update(e *entity.Account) error {
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

// Upsert mocks base method.
func (m *MockRepository) Upsert(e *entity.Account) (id.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", e)
	ret0, _ := ret[0].(id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRepositoryMockRecorder) Upsert(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRepository)(nil).Upsert), e)
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

// CreateAccount mocks base method.
func (m *MockUseCase) CreateAccount(tenantID id.ID, cognitoID, username, first_name, last_name, phone, email string, at entity.AccountType, as entity.AccountStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", tenantID, cognitoID, username, first_name, last_name, phone, email, at, as)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockUseCaseMockRecorder) CreateAccount(tenantID, cognitoID, username, first_name, last_name, phone, email, at, as interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockUseCase)(nil).CreateAccount), tenantID, cognitoID, username, first_name, last_name, phone, email, at, as)
}

// DeleteAccount mocks base method.
func (m *MockUseCase) DeleteAccount(tenantID, accountID id.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", tenantID, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockUseCaseMockRecorder) DeleteAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockUseCase)(nil).DeleteAccount), tenantID, accountID)
}

// DeleteAccountByName mocks base method.
func (m *MockUseCase) DeleteAccountByName(tenantID id.ID, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccountByName", tenantID, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccountByName indicates an expected call of DeleteAccountByName.
func (mr *MockUseCaseMockRecorder) DeleteAccountByName(tenantID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccountByName", reflect.TypeOf((*MockUseCase)(nil).DeleteAccountByName), tenantID, username)
}

// GetAccount mocks base method.
func (m *MockUseCase) GetAccount(tenantID, accountID id.ID) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", tenantID, accountID)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockUseCaseMockRecorder) GetAccount(tenantID, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockUseCase)(nil).GetAccount), tenantID, accountID)
}

// GetAccountByEmail mocks base method.
func (m *MockUseCase) GetAccountByEmail(tenantID id.ID, email string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByEmail", tenantID, email)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByEmail indicates an expected call of GetAccountByEmail.
func (mr *MockUseCaseMockRecorder) GetAccountByEmail(tenantID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByEmail", reflect.TypeOf((*MockUseCase)(nil).GetAccountByEmail), tenantID, email)
}

// GetAccountByName mocks base method.
func (m *MockUseCase) GetAccountByName(tenantID id.ID, username string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByName", tenantID, username)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByName indicates an expected call of GetAccountByName.
func (mr *MockUseCaseMockRecorder) GetAccountByName(tenantID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByName", reflect.TypeOf((*MockUseCase)(nil).GetAccountByName), tenantID, username)
}

// GetCount mocks base method.
func (m *MockUseCase) GetCount(tenantId id.ID) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", tenantId)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetCount indicates an expected call of GetCount.
func (mr *MockUseCaseMockRecorder) GetCount(tenantId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockUseCase)(nil).GetCount), tenantId)
}

// ListAccounts mocks base method.
func (m *MockUseCase) ListAccounts(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", tenantID, page, limit, at)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockUseCaseMockRecorder) ListAccounts(tenantID, page, limit, at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockUseCase)(nil).ListAccounts), tenantID, page, limit, at)
}

// SearchAccounts mocks base method.
func (m *MockUseCase) SearchAccounts(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchAccounts", tenantID, query, page, limit, at)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAccounts indicates an expected call of SearchAccounts.
func (mr *MockUseCaseMockRecorder) SearchAccounts(tenantID, query, page, limit, at interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAccounts", reflect.TypeOf((*MockUseCase)(nil).SearchAccounts), tenantID, query, page, limit, at)
}

// UpdateAccount mocks base method.
func (m *MockUseCase) UpdateAccount(e *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockUseCaseMockRecorder) UpdateAccount(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockUseCase)(nil).UpdateAccount), e)
}

// UpsertAccount mocks base method.
func (m *MockUseCase) UpsertAccount(e *entity.Account) (id.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertAccount", e)
	ret0, _ := ret[0].(id.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertAccount indicates an expected call of UpsertAccount.
func (mr *MockUseCaseMockRecorder) UpsertAccount(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertAccount", reflect.TypeOf((*MockUseCase)(nil).UpsertAccount), e)
}
