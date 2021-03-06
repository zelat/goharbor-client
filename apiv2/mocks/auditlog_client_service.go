// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	auditlog "github.com/zelat/goharbor-client/apiv2/internal/api/client/auditlog"

	runtime "github.com/go-openapi/runtime"
)

// MockAuditlogClientService is an autogenerated mock type for the ClientService type
type MockAuditlogClientService struct {
	mock.Mock
}

// ListAuditLogs provides a mock function with given fields: params, authInfo, opts
func (_m *MockAuditlogClientService) ListAuditLogs(params *auditlog.ListAuditLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...auditlog.ClientOption) (*auditlog.ListAuditLogsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *auditlog.ListAuditLogsOK
	if rf, ok := ret.Get(0).(func(*auditlog.ListAuditLogsParams, runtime.ClientAuthInfoWriter, ...auditlog.ClientOption) *auditlog.ListAuditLogsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auditlog.ListAuditLogsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*auditlog.ListAuditLogsParams, runtime.ClientAuthInfoWriter, ...auditlog.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockAuditlogClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}
