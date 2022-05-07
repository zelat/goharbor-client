// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	ping "github.com/zelat/goharbor-client/v4/apiv2/internal/api/client/ping"
	mock "github.com/stretchr/testify/mock"
)

// MockPingClientService is an autogenerated mock type for the ClientService type
type MockPingClientService struct {
	mock.Mock
}

// GetPing provides a mock function with given fields: params, authInfo, opts
func (_m *MockPingClientService) GetPing(params *ping.GetPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ping.ClientOption) (*ping.GetPingOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ping.GetPingOK
	if rf, ok := ret.Get(0).(func(*ping.GetPingParams, runtime.ClientAuthInfoWriter, ...ping.ClientOption) *ping.GetPingOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ping.GetPingOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ping.GetPingParams, runtime.ClientAuthInfoWriter, ...ping.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockPingClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
