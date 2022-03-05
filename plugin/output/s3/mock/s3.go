// Code generated by MockGen. DO NOT EDIT.
// Source: plugin/output/s3/s3.go

// Package mock_s3 is a generated GoMock package.
package mock_s3

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	minio "github.com/minio/minio-go"
)

// MockObjectStoreClient is a mock of ObjectStoreClient interface.
type MockObjectStoreClient struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStoreClientMockRecorder
}

// MockObjectStoreClientMockRecorder is the mock recorder for MockObjectStoreClient.
type MockObjectStoreClientMockRecorder struct {
	mock *MockObjectStoreClient
}

// NewMockObjectStoreClient creates a new mock instance.
func NewMockObjectStoreClient(ctrl *gomock.Controller) *MockObjectStoreClient {
	mock := &MockObjectStoreClient{ctrl: ctrl}
	mock.recorder = &MockObjectStoreClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStoreClient) EXPECT() *MockObjectStoreClientMockRecorder {
	return m.recorder
}

// BucketExists mocks base method.
func (m *MockObjectStoreClient) BucketExists(bucketName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BucketExists", bucketName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BucketExists indicates an expected call of BucketExists.
func (mr *MockObjectStoreClientMockRecorder) BucketExists(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BucketExists", reflect.TypeOf((*MockObjectStoreClient)(nil).BucketExists), bucketName)
}

// FPutObject mocks base method.
func (m *MockObjectStoreClient) FPutObject(bucketName, objectName, filePath string, opts minio.PutObjectOptions) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FPutObject", bucketName, objectName, filePath, opts)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FPutObject indicates an expected call of FPutObject.
func (mr *MockObjectStoreClientMockRecorder) FPutObject(bucketName, objectName, filePath, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FPutObject", reflect.TypeOf((*MockObjectStoreClient)(nil).FPutObject), bucketName, objectName, filePath, opts)
}

// MakeBucket mocks base method.
func (m *MockObjectStoreClient) MakeBucket(bucketName, location string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeBucket", bucketName, location)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeBucket indicates an expected call of MakeBucket.
func (mr *MockObjectStoreClientMockRecorder) MakeBucket(bucketName, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeBucket", reflect.TypeOf((*MockObjectStoreClient)(nil).MakeBucket), bucketName, location)
}

// Mockcompressor is a mock of compressor interface.
type Mockcompressor struct {
	ctrl     *gomock.Controller
	recorder *MockcompressorMockRecorder
}

// MockcompressorMockRecorder is the mock recorder for Mockcompressor.
type MockcompressorMockRecorder struct {
	mock *Mockcompressor
}

// NewMockcompressor creates a new mock instance.
func NewMockcompressor(ctrl *gomock.Controller) *Mockcompressor {
	mock := &Mockcompressor{ctrl: ctrl}
	mock.recorder = &MockcompressorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcompressor) EXPECT() *MockcompressorMockRecorder {
	return m.recorder
}

// compress mocks base method.
func (m *Mockcompressor) compress(archiveName, fileName string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "compress", archiveName, fileName)
}

// compress indicates an expected call of compress.
func (mr *MockcompressorMockRecorder) compress(archiveName, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "compress", reflect.TypeOf((*Mockcompressor)(nil).compress), archiveName, fileName)
}

// getExtension mocks base method.
func (m *Mockcompressor) getExtension() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getExtension")
	ret0, _ := ret[0].(string)
	return ret0
}

// getExtension indicates an expected call of getExtension.
func (mr *MockcompressorMockRecorder) getExtension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getExtension", reflect.TypeOf((*Mockcompressor)(nil).getExtension))
}

// getName mocks base method.
func (m *Mockcompressor) getName(fileName string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getName", fileName)
	ret0, _ := ret[0].(string)
	return ret0
}

// getName indicates an expected call of getName.
func (mr *MockcompressorMockRecorder) getName(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getName", reflect.TypeOf((*Mockcompressor)(nil).getName), fileName)
}

// getObjectOptions mocks base method.
func (m *Mockcompressor) getObjectOptions() minio.PutObjectOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getObjectOptions")
	ret0, _ := ret[0].(minio.PutObjectOptions)
	return ret0
}

// getObjectOptions indicates an expected call of getObjectOptions.
func (mr *MockcompressorMockRecorder) getObjectOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getObjectOptions", reflect.TypeOf((*Mockcompressor)(nil).getObjectOptions))
}