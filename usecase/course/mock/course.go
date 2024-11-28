// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/course/interface.go

// Package mock_course is a generated GoMock package.
package mock_course

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

// Get mocks base method.
func (m *MockReader) Get(id entity.ID) (*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReader)(nil).Get), id)
}

// GetCount mocks base method.
func (m *MockReader) GetCount(id entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockReaderMockRecorder) GetCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockReader)(nil).GetCount), id)
}

// List mocks base method.
func (m *MockReader) List(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tenantID, page, limit)
	ret0, _ := ret[0].([]*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockReaderMockRecorder) List(tenantID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockReader)(nil).List), tenantID, page, limit)
}

// Search mocks base method.
func (m *MockReader) Search(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", tenantID, query, page, limit)
	ret0, _ := ret[0].([]*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockReaderMockRecorder) Search(tenantID, query, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockReader)(nil).Search), tenantID, query, page, limit)
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
func (m *MockWriter) Create(e *entity.Course) (entity.ID, error) {
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
func (m *MockWriter) Update(e *entity.Course) error {
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
func (m *MockRepository) Create(e *entity.Course) (entity.ID, error) {
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
func (m *MockRepository) Get(id entity.ID) (*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepository)(nil).Get), id)
}

// GetCount mocks base method.
func (m *MockRepository) GetCount(id entity.ID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockRepositoryMockRecorder) GetCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockRepository)(nil).GetCount), id)
}

// List mocks base method.
func (m *MockRepository) List(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tenantID, page, limit)
	ret0, _ := ret[0].([]*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockRepositoryMockRecorder) List(tenantID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepository)(nil).List), tenantID, page, limit)
}

// Search mocks base method.
func (m *MockRepository) Search(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", tenantID, query, page, limit)
	ret0, _ := ret[0].([]*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRepositoryMockRecorder) Search(tenantID, query, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepository)(nil).Search), tenantID, query, page, limit)
}

// Update mocks base method.
func (m *MockRepository) Update(e *entity.Course) error {
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

// CreateCourse mocks base method.
func (m *MockUseCase) CreateCourse(tenantID entity.ID, extID *string, centerID, productID entity.ID, name, notes, timezone string, address entity.CourseAddress, status entity.CourseStatus, mode entity.CourseMode, maxAttendees, numAttendees int32) (entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCourse", tenantID, extID, centerID, productID, name, notes, timezone, address, status, mode, maxAttendees, numAttendees)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCourse indicates an expected call of CreateCourse.
func (mr *MockUseCaseMockRecorder) CreateCourse(tenantID, extID, centerID, productID, name, notes, timezone, address, status, mode, maxAttendees, numAttendees interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCourse", reflect.TypeOf((*MockUseCase)(nil).CreateCourse), tenantID, extID, centerID, productID, name, notes, timezone, address, status, mode, maxAttendees, numAttendees)
}

// DeleteCourse mocks base method.
func (m *MockUseCase) DeleteCourse(id entity.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCourse", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCourse indicates an expected call of DeleteCourse.
func (mr *MockUseCaseMockRecorder) DeleteCourse(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCourse", reflect.TypeOf((*MockUseCase)(nil).DeleteCourse), id)
}

// GetCount mocks base method.
func (m *MockUseCase) GetCount(id entity.ID) int {
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

// GetCourse mocks base method.
func (m *MockUseCase) GetCourse(id entity.ID) (*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourse", id)
	ret0, _ := ret[0].(*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourse indicates an expected call of GetCourse.
func (mr *MockUseCaseMockRecorder) GetCourse(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourse", reflect.TypeOf((*MockUseCase)(nil).GetCourse), id)
}

// ListCourses mocks base method.
func (m *MockUseCase) ListCourses(tenantID entity.ID, page, limit int) ([]*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCourses", tenantID, page, limit)
	ret0, _ := ret[0].([]*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCourses indicates an expected call of ListCourses.
func (mr *MockUseCaseMockRecorder) ListCourses(tenantID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCourses", reflect.TypeOf((*MockUseCase)(nil).ListCourses), tenantID, page, limit)
}

// SearchCourses mocks base method.
func (m *MockUseCase) SearchCourses(tenantID entity.ID, query string, page, limit int) ([]*entity.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCourses", tenantID, query, page, limit)
	ret0, _ := ret[0].([]*entity.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCourses indicates an expected call of SearchCourses.
func (mr *MockUseCaseMockRecorder) SearchCourses(tenantID, query, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCourses", reflect.TypeOf((*MockUseCase)(nil).SearchCourses), tenantID, query, page, limit)
}

// UpdateCourse mocks base method.
func (m *MockUseCase) UpdateCourse(e *entity.Course) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourse", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCourse indicates an expected call of UpdateCourse.
func (mr *MockUseCaseMockRecorder) UpdateCourse(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourse", reflect.TypeOf((*MockUseCase)(nil).UpdateCourse), e)
}
