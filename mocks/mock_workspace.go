// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/archer/workspace.go

// Package mocks is a generated GoMock package.
package mocks

import (
	archer "github.com/aws/amazon-ecs-cli-v2/internal/pkg/archer"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWorkspace is a mock of Workspace interface
type MockWorkspace struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceMockRecorder
}

// MockWorkspaceMockRecorder is the mock recorder for MockWorkspace
type MockWorkspaceMockRecorder struct {
	mock *MockWorkspace
}

// NewMockWorkspace creates a new mock instance
func NewMockWorkspace(ctrl *gomock.Controller) *MockWorkspace {
	mock := &MockWorkspace{ctrl: ctrl}
	mock.recorder = &MockWorkspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkspace) EXPECT() *MockWorkspaceMockRecorder {
	return m.recorder
}

// WriteFile mocks base method
func (m *MockWorkspace) WriteFile(blob []byte, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", blob, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockWorkspaceMockRecorder) WriteFile(blob, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockWorkspace)(nil).WriteFile), blob, filename)
}

// ReadFile mocks base method
func (m *MockWorkspace) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockWorkspaceMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockWorkspace)(nil).ReadFile), filename)
}

// ListManifestFiles mocks base method
func (m *MockWorkspace) ListManifestFiles() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListManifestFiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListManifestFiles indicates an expected call of ListManifestFiles
func (mr *MockWorkspaceMockRecorder) ListManifestFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManifestFiles", reflect.TypeOf((*MockWorkspace)(nil).ListManifestFiles))
}

// AppManifestFileName mocks base method
func (m *MockWorkspace) AppManifestFileName(appName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppManifestFileName", appName)
	ret0, _ := ret[0].(string)
	return ret0
}

// AppManifestFileName indicates an expected call of AppManifestFileName
func (mr *MockWorkspaceMockRecorder) AppManifestFileName(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppManifestFileName", reflect.TypeOf((*MockWorkspace)(nil).AppManifestFileName), appName)
}

// DeleteFile mocks base method
func (m *MockWorkspace) DeleteFile(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile
func (mr *MockWorkspaceMockRecorder) DeleteFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockWorkspace)(nil).DeleteFile), name)
}

// Create mocks base method
func (m *MockWorkspace) Create(projectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", projectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockWorkspaceMockRecorder) Create(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWorkspace)(nil).Create), projectName)
}

// Summary mocks base method
func (m *MockWorkspace) Summary() (*archer.WorkspaceSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Summary")
	ret0, _ := ret[0].(*archer.WorkspaceSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Summary indicates an expected call of Summary
func (mr *MockWorkspaceMockRecorder) Summary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Summary", reflect.TypeOf((*MockWorkspace)(nil).Summary))
}

// Apps mocks base method
func (m *MockWorkspace) Apps() ([]archer.Manifest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apps")
	ret0, _ := ret[0].([]archer.Manifest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apps indicates an expected call of Apps
func (mr *MockWorkspaceMockRecorder) Apps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apps", reflect.TypeOf((*MockWorkspace)(nil).Apps))
}

// MockManifestIO is a mock of ManifestIO interface
type MockManifestIO struct {
	ctrl     *gomock.Controller
	recorder *MockManifestIOMockRecorder
}

// MockManifestIOMockRecorder is the mock recorder for MockManifestIO
type MockManifestIOMockRecorder struct {
	mock *MockManifestIO
}

// NewMockManifestIO creates a new mock instance
func NewMockManifestIO(ctrl *gomock.Controller) *MockManifestIO {
	mock := &MockManifestIO{ctrl: ctrl}
	mock.recorder = &MockManifestIOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManifestIO) EXPECT() *MockManifestIOMockRecorder {
	return m.recorder
}

// WriteFile mocks base method
func (m *MockManifestIO) WriteFile(blob []byte, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", blob, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockManifestIOMockRecorder) WriteFile(blob, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockManifestIO)(nil).WriteFile), blob, filename)
}

// ReadFile mocks base method
func (m *MockManifestIO) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockManifestIOMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockManifestIO)(nil).ReadFile), filename)
}

// ListManifestFiles mocks base method
func (m *MockManifestIO) ListManifestFiles() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListManifestFiles")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListManifestFiles indicates an expected call of ListManifestFiles
func (mr *MockManifestIOMockRecorder) ListManifestFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListManifestFiles", reflect.TypeOf((*MockManifestIO)(nil).ListManifestFiles))
}

// AppManifestFileName mocks base method
func (m *MockManifestIO) AppManifestFileName(appName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppManifestFileName", appName)
	ret0, _ := ret[0].(string)
	return ret0
}

// AppManifestFileName indicates an expected call of AppManifestFileName
func (mr *MockManifestIOMockRecorder) AppManifestFileName(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppManifestFileName", reflect.TypeOf((*MockManifestIO)(nil).AppManifestFileName), appName)
}

// DeleteFile mocks base method
func (m *MockManifestIO) DeleteFile(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile
func (mr *MockManifestIOMockRecorder) DeleteFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockManifestIO)(nil).DeleteFile), name)
}

// MockWorkspaceFileReadWriter is a mock of WorkspaceFileReadWriter interface
type MockWorkspaceFileReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceFileReadWriterMockRecorder
}

// MockWorkspaceFileReadWriterMockRecorder is the mock recorder for MockWorkspaceFileReadWriter
type MockWorkspaceFileReadWriterMockRecorder struct {
	mock *MockWorkspaceFileReadWriter
}

// NewMockWorkspaceFileReadWriter creates a new mock instance
func NewMockWorkspaceFileReadWriter(ctrl *gomock.Controller) *MockWorkspaceFileReadWriter {
	mock := &MockWorkspaceFileReadWriter{ctrl: ctrl}
	mock.recorder = &MockWorkspaceFileReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkspaceFileReadWriter) EXPECT() *MockWorkspaceFileReadWriterMockRecorder {
	return m.recorder
}

// WriteFile mocks base method
func (m *MockWorkspaceFileReadWriter) WriteFile(blob []byte, filename string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", blob, filename)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockWorkspaceFileReadWriterMockRecorder) WriteFile(blob, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockWorkspaceFileReadWriter)(nil).WriteFile), blob, filename)
}

// ReadFile mocks base method
func (m *MockWorkspaceFileReadWriter) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockWorkspaceFileReadWriterMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockWorkspaceFileReadWriter)(nil).ReadFile), filename)
}
