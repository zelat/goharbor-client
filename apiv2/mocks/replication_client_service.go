// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"
	replication "github.com/zelat/goharbor-client/apiv2/internal/api/client/replication"
)

// MockReplicationClientService is an autogenerated mock type for the ClientService type
type MockReplicationClientService struct {
	mock.Mock
}

// GetReplicationExecution provides a mock function with given fields: params, authInfo, opts
func (_m *MockReplicationClientService) GetReplicationExecution(params *replication.GetReplicationExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...replication.ClientOption) (*replication.GetReplicationExecutionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *replication.GetReplicationExecutionOK
	if rf, ok := ret.Get(0).(func(*replication.GetReplicationExecutionParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) *replication.GetReplicationExecutionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.GetReplicationExecutionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.GetReplicationExecutionParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReplicationLog provides a mock function with given fields: params, authInfo, opts
func (_m *MockReplicationClientService) GetReplicationLog(params *replication.GetReplicationLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...replication.ClientOption) (*replication.GetReplicationLogOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *replication.GetReplicationLogOK
	if rf, ok := ret.Get(0).(func(*replication.GetReplicationLogParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) *replication.GetReplicationLogOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.GetReplicationLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.GetReplicationLogParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicationExecutions provides a mock function with given fields: params, authInfo, opts
func (_m *MockReplicationClientService) ListReplicationExecutions(params *replication.ListReplicationExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...replication.ClientOption) (*replication.ListReplicationExecutionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *replication.ListReplicationExecutionsOK
	if rf, ok := ret.Get(0).(func(*replication.ListReplicationExecutionsParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) *replication.ListReplicationExecutionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.ListReplicationExecutionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.ListReplicationExecutionsParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListReplicationTasks provides a mock function with given fields: params, authInfo, opts
func (_m *MockReplicationClientService) ListReplicationTasks(params *replication.ListReplicationTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...replication.ClientOption) (*replication.ListReplicationTasksOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *replication.ListReplicationTasksOK
	if rf, ok := ret.Get(0).(func(*replication.ListReplicationTasksParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) *replication.ListReplicationTasksOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.ListReplicationTasksOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.ListReplicationTasksParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockReplicationClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// StartReplication provides a mock function with given fields: params, authInfo, opts
func (_m *MockReplicationClientService) StartReplication(params *replication.StartReplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...replication.ClientOption) (*replication.StartReplicationCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *replication.StartReplicationCreated
	if rf, ok := ret.Get(0).(func(*replication.StartReplicationParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) *replication.StartReplicationCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.StartReplicationCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.StartReplicationParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StopReplication provides a mock function with given fields: params, authInfo, opts
func (_m *MockReplicationClientService) StopReplication(params *replication.StopReplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...replication.ClientOption) (*replication.StopReplicationOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *replication.StopReplicationOK
	if rf, ok := ret.Get(0).(func(*replication.StopReplicationParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) *replication.StopReplicationOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*replication.StopReplicationOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*replication.StopReplicationParams, runtime.ClientAuthInfoWriter, ...replication.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
