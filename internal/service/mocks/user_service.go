// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	entity "github.com/sashaaro/gophkeeper/internal/entity"
)

// MockPasswordHasher is a mock of PasswordHasher interface.
type MockPasswordHasher struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordHasherMockRecorder
}

// MockPasswordHasherMockRecorder is the mock recorder for MockPasswordHasher.
type MockPasswordHasherMockRecorder struct {
	mock *MockPasswordHasher
}

// NewMockPasswordHasher creates a new mock instance.
func NewMockPasswordHasher(ctrl *gomock.Controller) *MockPasswordHasher {
	mock := &MockPasswordHasher{ctrl: ctrl}
	mock.recorder = &MockPasswordHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordHasher) EXPECT() *MockPasswordHasherMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockPasswordHasher) Hash(pwd string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", pwd)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash.
func (mr *MockPasswordHasherMockRecorder) Hash(pwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockPasswordHasher)(nil).Hash), pwd)
}

// IsEqual mocks base method.
func (m *MockPasswordHasher) IsEqual(hashed, check string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEqual", hashed, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsEqual indicates an expected call of IsEqual.
func (mr *MockPasswordHasherMockRecorder) IsEqual(hashed, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEqual", reflect.TypeOf((*MockPasswordHasher)(nil).IsEqual), hashed, check)
}

// MockUserCreator is a mock of UserCreator interface.
type MockUserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockUserCreatorMockRecorder
}

// MockUserCreatorMockRecorder is the mock recorder for MockUserCreator.
type MockUserCreatorMockRecorder struct {
	mock *MockUserCreator
}

// NewMockUserCreator creates a new mock instance.
func NewMockUserCreator(ctrl *gomock.Controller) *MockUserCreator {
	mock := &MockUserCreator{ctrl: ctrl}
	mock.recorder = &MockUserCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserCreator) EXPECT() *MockUserCreatorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserCreator) Create(ctx context.Context, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserCreatorMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserCreator)(nil).Create), ctx, m)
}

// MockUserGetter is a mock of UserGetter interface.
type MockUserGetter struct {
	ctrl     *gomock.Controller
	recorder *MockUserGetterMockRecorder
}

// MockUserGetterMockRecorder is the mock recorder for MockUserGetter.
type MockUserGetterMockRecorder struct {
	mock *MockUserGetter
}

// NewMockUserGetter creates a new mock instance.
func NewMockUserGetter(ctrl *gomock.Controller) *MockUserGetter {
	mock := &MockUserGetter{ctrl: ctrl}
	mock.recorder = &MockUserGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserGetter) EXPECT() *MockUserGetterMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m_2 *MockUserGetter) Get(ctx context.Context, id uuid.UUID, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Get", ctx, id, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockUserGetterMockRecorder) Get(ctx, id, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserGetter)(nil).Get), ctx, id, m)
}

// GetByLogin mocks base method.
func (m_2 *MockUserGetter) GetByLogin(ctx context.Context, login string, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetByLogin", ctx, login, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByLogin indicates an expected call of GetByLogin.
func (mr *MockUserGetterMockRecorder) GetByLogin(ctx, login, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockUserGetter)(nil).GetByLogin), ctx, login, m)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserRepository) Create(ctx context.Context, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, m)
}

// Get mocks base method.
func (m_2 *MockUserRepository) Get(ctx context.Context, id uuid.UUID, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Get", ctx, id, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockUserRepositoryMockRecorder) Get(ctx, id, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRepository)(nil).Get), ctx, id, m)
}

// GetByLogin mocks base method.
func (m_2 *MockUserRepository) GetByLogin(ctx context.Context, login string, m *entity.User) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetByLogin", ctx, login, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByLogin indicates an expected call of GetByLogin.
func (mr *MockUserRepositoryMockRecorder) GetByLogin(ctx, login, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockUserRepository)(nil).GetByLogin), ctx, login, m)
}
