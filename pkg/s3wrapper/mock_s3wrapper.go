// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package s3wrapper is a generated GoMock package.
package s3wrapper

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method
func (m *MockAPI) CreateBucket() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket
func (mr *MockAPIMockRecorder) CreateBucket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockAPI)(nil).CreateBucket))
}

// Upload mocks base method
func (m *MockAPI) Upload(ctx context.Context, data []byte, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", ctx, data, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockAPIMockRecorder) Upload(ctx, data, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockAPI)(nil).Upload), ctx, data, objectName)
}

// Download mocks base method
func (m *MockAPI) Download(ctx context.Context, objectName string) (io.ReadCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", ctx, objectName)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Download indicates an expected call of Download
func (mr *MockAPIMockRecorder) Download(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockAPI)(nil).Download), ctx, objectName)
}

// DoesObjectExist mocks base method
func (m *MockAPI) DoesObjectExist(ctx context.Context, objectName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesObjectExist", ctx, objectName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesObjectExist indicates an expected call of DoesObjectExist
func (mr *MockAPIMockRecorder) DoesObjectExist(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesObjectExist", reflect.TypeOf((*MockAPI)(nil).DoesObjectExist), ctx, objectName)
}

// DeleteObject mocks base method
func (m *MockAPI) DeleteObject(ctx context.Context, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockAPIMockRecorder) DeleteObject(ctx, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockAPI)(nil).DeleteObject), ctx, objectName)
}

// UpdateObjectTag mocks base method
func (m *MockAPI) UpdateObjectTag(ctx context.Context, objectName, key, value string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectTag", ctx, objectName, key, value)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateObjectTag indicates an expected call of UpdateObjectTag
func (mr *MockAPIMockRecorder) UpdateObjectTag(ctx, objectName, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectTag", reflect.TypeOf((*MockAPI)(nil).UpdateObjectTag), ctx, objectName, key, value)
}