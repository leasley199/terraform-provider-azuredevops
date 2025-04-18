// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/azure-devops-go-api/azuredevops/v7/policy (interfaces: Client)

// Package azdosdkmocks is a generated GoMock package.
package azdosdkmocks

import (
	context "context"
	reflect "reflect"

	policy "github.com/microsoft/azure-devops-go-api/azuredevops/v7/policy"
	gomock "go.uber.org/mock/gomock"
)

// MockPolicyClient is a mock of Client interface.
type MockPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyClientMockRecorder
	isgomock struct{}
}

// MockPolicyClientMockRecorder is the mock recorder for MockPolicyClient.
type MockPolicyClientMockRecorder struct {
	mock *MockPolicyClient
}

// NewMockPolicyClient creates a new mock instance.
func NewMockPolicyClient(ctrl *gomock.Controller) *MockPolicyClient {
	mock := &MockPolicyClient{ctrl: ctrl}
	mock.recorder = &MockPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPolicyClient) EXPECT() *MockPolicyClientMockRecorder {
	return m.recorder
}

// CreatePolicyConfiguration mocks base method.
func (m *MockPolicyClient) CreatePolicyConfiguration(arg0 context.Context, arg1 policy.CreatePolicyConfigurationArgs) (*policy.PolicyConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePolicyConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePolicyConfiguration indicates an expected call of CreatePolicyConfiguration.
func (mr *MockPolicyClientMockRecorder) CreatePolicyConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePolicyConfiguration", reflect.TypeOf((*MockPolicyClient)(nil).CreatePolicyConfiguration), arg0, arg1)
}

// DeletePolicyConfiguration mocks base method.
func (m *MockPolicyClient) DeletePolicyConfiguration(arg0 context.Context, arg1 policy.DeletePolicyConfigurationArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicyConfiguration", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicyConfiguration indicates an expected call of DeletePolicyConfiguration.
func (mr *MockPolicyClientMockRecorder) DeletePolicyConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicyConfiguration", reflect.TypeOf((*MockPolicyClient)(nil).DeletePolicyConfiguration), arg0, arg1)
}

// GetPolicyConfiguration mocks base method.
func (m *MockPolicyClient) GetPolicyConfiguration(arg0 context.Context, arg1 policy.GetPolicyConfigurationArgs) (*policy.PolicyConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyConfiguration indicates an expected call of GetPolicyConfiguration.
func (mr *MockPolicyClientMockRecorder) GetPolicyConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyConfiguration", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyConfiguration), arg0, arg1)
}

// GetPolicyConfigurationRevision mocks base method.
func (m *MockPolicyClient) GetPolicyConfigurationRevision(arg0 context.Context, arg1 policy.GetPolicyConfigurationRevisionArgs) (*policy.PolicyConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyConfigurationRevision", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyConfigurationRevision indicates an expected call of GetPolicyConfigurationRevision.
func (mr *MockPolicyClientMockRecorder) GetPolicyConfigurationRevision(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyConfigurationRevision", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyConfigurationRevision), arg0, arg1)
}

// GetPolicyConfigurationRevisions mocks base method.
func (m *MockPolicyClient) GetPolicyConfigurationRevisions(arg0 context.Context, arg1 policy.GetPolicyConfigurationRevisionsArgs) (*[]policy.PolicyConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyConfigurationRevisions", arg0, arg1)
	ret0, _ := ret[0].(*[]policy.PolicyConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyConfigurationRevisions indicates an expected call of GetPolicyConfigurationRevisions.
func (mr *MockPolicyClientMockRecorder) GetPolicyConfigurationRevisions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyConfigurationRevisions", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyConfigurationRevisions), arg0, arg1)
}

// GetPolicyConfigurations mocks base method.
func (m *MockPolicyClient) GetPolicyConfigurations(arg0 context.Context, arg1 policy.GetPolicyConfigurationsArgs) (*policy.GetPolicyConfigurationsResponseValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyConfigurations", arg0, arg1)
	ret0, _ := ret[0].(*policy.GetPolicyConfigurationsResponseValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyConfigurations indicates an expected call of GetPolicyConfigurations.
func (mr *MockPolicyClientMockRecorder) GetPolicyConfigurations(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyConfigurations", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyConfigurations), arg0, arg1)
}

// GetPolicyEvaluation mocks base method.
func (m *MockPolicyClient) GetPolicyEvaluation(arg0 context.Context, arg1 policy.GetPolicyEvaluationArgs) (*policy.PolicyEvaluationRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyEvaluation", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyEvaluationRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyEvaluation indicates an expected call of GetPolicyEvaluation.
func (mr *MockPolicyClientMockRecorder) GetPolicyEvaluation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyEvaluation", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyEvaluation), arg0, arg1)
}

// GetPolicyEvaluations mocks base method.
func (m *MockPolicyClient) GetPolicyEvaluations(arg0 context.Context, arg1 policy.GetPolicyEvaluationsArgs) (*[]policy.PolicyEvaluationRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyEvaluations", arg0, arg1)
	ret0, _ := ret[0].(*[]policy.PolicyEvaluationRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyEvaluations indicates an expected call of GetPolicyEvaluations.
func (mr *MockPolicyClientMockRecorder) GetPolicyEvaluations(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyEvaluations", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyEvaluations), arg0, arg1)
}

// GetPolicyType mocks base method.
func (m *MockPolicyClient) GetPolicyType(arg0 context.Context, arg1 policy.GetPolicyTypeArgs) (*policy.PolicyType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyType", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyType indicates an expected call of GetPolicyType.
func (mr *MockPolicyClientMockRecorder) GetPolicyType(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyType", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyType), arg0, arg1)
}

// GetPolicyTypes mocks base method.
func (m *MockPolicyClient) GetPolicyTypes(arg0 context.Context, arg1 policy.GetPolicyTypesArgs) (*[]policy.PolicyType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyTypes", arg0, arg1)
	ret0, _ := ret[0].(*[]policy.PolicyType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyTypes indicates an expected call of GetPolicyTypes.
func (mr *MockPolicyClientMockRecorder) GetPolicyTypes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyTypes", reflect.TypeOf((*MockPolicyClient)(nil).GetPolicyTypes), arg0, arg1)
}

// RequeuePolicyEvaluation mocks base method.
func (m *MockPolicyClient) RequeuePolicyEvaluation(arg0 context.Context, arg1 policy.RequeuePolicyEvaluationArgs) (*policy.PolicyEvaluationRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequeuePolicyEvaluation", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyEvaluationRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequeuePolicyEvaluation indicates an expected call of RequeuePolicyEvaluation.
func (mr *MockPolicyClientMockRecorder) RequeuePolicyEvaluation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequeuePolicyEvaluation", reflect.TypeOf((*MockPolicyClient)(nil).RequeuePolicyEvaluation), arg0, arg1)
}

// UpdatePolicyConfiguration mocks base method.
func (m *MockPolicyClient) UpdatePolicyConfiguration(arg0 context.Context, arg1 policy.UpdatePolicyConfigurationArgs) (*policy.PolicyConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePolicyConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*policy.PolicyConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePolicyConfiguration indicates an expected call of UpdatePolicyConfiguration.
func (mr *MockPolicyClientMockRecorder) UpdatePolicyConfiguration(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePolicyConfiguration", reflect.TypeOf((*MockPolicyClient)(nil).UpdatePolicyConfiguration), arg0, arg1)
}
