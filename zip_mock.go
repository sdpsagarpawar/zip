package zip

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockZipOperations is a mock of ZipOperations interface.
type MockZipOperations struct {
	ctrl     *gomock.Controller
	recorder *MockZipOperationsMockRecorder
}

// MockZipOperationsMockRecorder is the mock recorder for MockZipOperations.
type MockZipOperationsMockRecorder struct {
	mock *MockZipOperations
}

// NewMockZipOperations creates a new mock instance.
func NewMockZipOperations(ctrl *gomock.Controller) *MockZipOperations {
	mock := &MockZipOperations{ctrl: ctrl}
	mock.recorder = &MockZipOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockZipOperations) EXPECT() *MockZipOperationsMockRecorder {
	return m.recorder
}

// ReadFiles mocks base method.
func (m *MockZipOperations) ReadFiles(ctx context.Context, dir string) (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFiles", ctx, dir)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFiles indicates an expected call of ReadFiles.
func (mr *MockZipOperationsMockRecorder) ReadFiles(ctx, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFiles", reflect.TypeOf((*MockZipOperations)(nil).ReadFiles), ctx, dir)
}

// Unzip mocks base method.
func (m *MockZipOperations) Unzip(ctx context.Context, bodyBytes []byte) (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unzip", ctx, bodyBytes)
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unzip indicates an expected call of Unzip.
func (mr *MockZipOperationsMockRecorder) Unzip(ctx, bodyBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unzip", reflect.TypeOf((*MockZipOperations)(nil).Unzip), ctx, bodyBytes)
}

// WriteZip mocks base method.
func (m *MockZipOperations) WriteZip(ctx context.Context, file []byte, zipName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteZip", ctx, file, zipName)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteZip indicates an expected call of WriteZip.
func (mr *MockZipOperationsMockRecorder) WriteZip(ctx, file, zipName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteZip", reflect.TypeOf((*MockZipOperations)(nil).WriteZip), ctx, file, zipName)
}

// Zip mocks base method.
func (m *MockZipOperations) Zip(ctx context.Context, files map[string][]byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Zip", ctx, files)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Zip indicates an expected call of Zip.
func (mr *MockZipOperationsMockRecorder) Zip(ctx, files interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Zip", reflect.TypeOf((*MockZipOperations)(nil).Zip), ctx, files)
}
