// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	telemetryv1alpha1 "github.com/kyma-project/telemetry-manager/apis/telemetry/v1alpha1"
	"github.com/stretchr/testify/mock"
)

// VariablesValidator is an autogenerated mock type for the VariablesValidator type
type VariablesValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: _a0, logPipeline, logPipelines
func (_m *VariablesValidator) Validate(logPipeline *telemetryv1alpha1.LogPipeline, logPipelines *telemetryv1alpha1.LogPipelineList) error {
	ret := _m.Called(logPipeline, logPipelines)

	var r0 error
	if rf, ok := ret.Get(0).(func(*telemetryv1alpha1.LogPipeline, *telemetryv1alpha1.LogPipelineList) error); ok {
		r0 = rf(logPipeline, logPipelines)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewVariablesValidator interface {
	mock.TestingT
	Cleanup(func())
}

// NewVariablesValidator creates a new instance of VariablesValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVariablesValidator(t mockConstructorTestingTNewVariablesValidator) *VariablesValidator {
	mock := &VariablesValidator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}