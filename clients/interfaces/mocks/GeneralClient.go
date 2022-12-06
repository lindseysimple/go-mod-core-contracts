// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"
)

// GeneralClient is an autogenerated mock type for the GeneralClient type
type GeneralClient struct {
	mock.Mock
}

// FetchConfiguration provides a mock function with given fields: ctx
func (_m *GeneralClient) FetchConfiguration(ctx context.Context) (common.ConfigResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.ConfigResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.ConfigResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.ConfigResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// FetchMetrics provides a mock function with given fields: ctx
func (_m *GeneralClient) FetchMetrics(ctx context.Context) (common.MetricsResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.MetricsResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.MetricsResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.MetricsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewGeneralClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralClient creates a new instance of GeneralClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralClient(t mockConstructorTestingTNewGeneralClient) *GeneralClient {
	mock := &GeneralClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
