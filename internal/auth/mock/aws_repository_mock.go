// Code generated by MockGen. DO NOT EDIT.
// Source: aws_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/mahfuz110244/api-mc/internal/models"
	minio "github.com/minio/minio-go/v7"
	reflect "reflect"
)

// MockAWSRepository is a mock of AWSRepository interface
type MockAWSRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAWSRepositoryMockRecorder
}

// MockAWSRepositoryMockRecorder is the mock recorder for MockAWSRepository
type MockAWSRepositoryMockRecorder struct {
	mock *MockAWSRepository
}

// NewMockAWSRepository creates a new mock instance
func NewMockAWSRepository(ctrl *gomock.Controller) *MockAWSRepository {
	mock := &MockAWSRepository{ctrl: ctrl}
	mock.recorder = &MockAWSRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAWSRepository) EXPECT() *MockAWSRepositoryMockRecorder {
	return m.recorder
}

// PutObject mocks base method
func (m *MockAWSRepository) PutObject(ctx context.Context, input models.UploadInput) (*minio.UploadInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutObject", ctx, input)
	ret0, _ := ret[0].(*minio.UploadInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject
func (mr *MockAWSRepositoryMockRecorder) PutObject(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockAWSRepository)(nil).PutObject), ctx, input)
}

// GetObject mocks base method
func (m *MockAWSRepository) GetObject(ctx context.Context, bucket, fileName string) (*minio.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObject", ctx, bucket, fileName)
	ret0, _ := ret[0].(*minio.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockAWSRepositoryMockRecorder) GetObject(ctx, bucket, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockAWSRepository)(nil).GetObject), ctx, bucket, fileName)
}

// RemoveObject mocks base method
func (m *MockAWSRepository) RemoveObject(ctx context.Context, bucket, fileName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveObject", ctx, bucket, fileName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveObject indicates an expected call of RemoveObject
func (mr *MockAWSRepositoryMockRecorder) RemoveObject(ctx, bucket, fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObject", reflect.TypeOf((*MockAWSRepository)(nil).RemoveObject), ctx, bucket, fileName)
}
