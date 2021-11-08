// Code generated by MockGen. DO NOT EDIT.
// Source: internal/sysaccess/common.go

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	os "os"
	reflect "reflect"
	time "time"

	sysutil "github.com/digitalocean/droplet-agent/internal/sysutil"
	gomock "github.com/golang/mock/gomock"
)

// MocksysManager is a mock of sysManager interface.
type MocksysManager struct {
	ctrl     *gomock.Controller
	recorder *MocksysManagerMockRecorder
}

// MocksysManagerMockRecorder is the mock recorder for MocksysManager.
type MocksysManagerMockRecorder struct {
	mock *MocksysManager
}

// NewMocksysManager creates a new mock instance.
func NewMocksysManager(ctrl *gomock.Controller) *MocksysManager {
	mock := &MocksysManager{ctrl: ctrl}
	mock.recorder = &MocksysManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksysManager) EXPECT() *MocksysManagerMockRecorder {
	return m.recorder
}

// CopyFileAttribute mocks base method.
func (m *MocksysManager) CopyFileAttribute(from, to string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFileAttribute", from, to)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyFileAttribute indicates an expected call of CopyFileAttribute.
func (mr *MocksysManagerMockRecorder) CopyFileAttribute(from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFileAttribute", reflect.TypeOf((*MocksysManager)(nil).CopyFileAttribute), from, to)
}

// CreateFileForWrite mocks base method.
func (m *MocksysManager) CreateFileForWrite(file string, user *sysutil.User, perm os.FileMode) (io.WriteCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFileForWrite", file, user, perm)
	ret0, _ := ret[0].(io.WriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFileForWrite indicates an expected call of CreateFileForWrite.
func (mr *MocksysManagerMockRecorder) CreateFileForWrite(file, user, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFileForWrite", reflect.TypeOf((*MocksysManager)(nil).CreateFileForWrite), file, user, perm)
}

// FileExists mocks base method.
func (m *MocksysManager) FileExists(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileExists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FileExists indicates an expected call of FileExists.
func (mr *MocksysManagerMockRecorder) FileExists(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileExists", reflect.TypeOf((*MocksysManager)(nil).FileExists), name)
}

// GetUserByName mocks base method.
func (m *MocksysManager) GetUserByName(username string) (*sysutil.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", username)
	ret0, _ := ret[0].(*sysutil.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MocksysManagerMockRecorder) GetUserByName(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MocksysManager)(nil).GetUserByName), username)
}

// MkDirIfNonExist mocks base method.
func (m *MocksysManager) MkDirIfNonExist(dir string, user *sysutil.User, perm os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkDirIfNonExist", dir, user, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkDirIfNonExist indicates an expected call of MkDirIfNonExist.
func (mr *MocksysManagerMockRecorder) MkDirIfNonExist(dir, user, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkDirIfNonExist", reflect.TypeOf((*MocksysManager)(nil).MkDirIfNonExist), dir, user, perm)
}

// ReadFile mocks base method.
func (m *MocksysManager) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MocksysManagerMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MocksysManager)(nil).ReadFile), filename)
}

// RemoveFile mocks base method.
func (m *MocksysManager) RemoveFile(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFile", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFile indicates an expected call of RemoveFile.
func (mr *MocksysManagerMockRecorder) RemoveFile(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFile", reflect.TypeOf((*MocksysManager)(nil).RemoveFile), name)
}

// RenameFile mocks base method.
func (m *MocksysManager) RenameFile(oldpath, newpath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameFile", oldpath, newpath)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameFile indicates an expected call of RenameFile.
func (mr *MocksysManagerMockRecorder) RenameFile(oldpath, newpath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameFile", reflect.TypeOf((*MocksysManager)(nil).RenameFile), oldpath, newpath)
}

// Sleep mocks base method.
func (m *MocksysManager) Sleep(d time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sleep", d)
}

// Sleep indicates an expected call of Sleep.
func (mr *MocksysManagerMockRecorder) Sleep(d interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MocksysManager)(nil).Sleep), d)
}
