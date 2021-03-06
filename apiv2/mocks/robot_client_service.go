// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"
	robot "github.com/zelat/goharbor-client/apiv2/internal/api/client/robot"
)

// MockRobotClientService is an autogenerated mock type for the ClientService type
type MockRobotClientService struct {
	mock.Mock
}

// CreateRobot provides a mock function with given fields: params, authInfo, opts
func (_m *MockRobotClientService) CreateRobot(params *robot.CreateRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...robot.ClientOption) (*robot.CreateRobotCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *robot.CreateRobotCreated
	if rf, ok := ret.Get(0).(func(*robot.CreateRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) *robot.CreateRobotCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.CreateRobotCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.CreateRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRobot provides a mock function with given fields: params, authInfo, opts
func (_m *MockRobotClientService) DeleteRobot(params *robot.DeleteRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...robot.ClientOption) (*robot.DeleteRobotOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *robot.DeleteRobotOK
	if rf, ok := ret.Get(0).(func(*robot.DeleteRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) *robot.DeleteRobotOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.DeleteRobotOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.DeleteRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRobotByID provides a mock function with given fields: params, authInfo, opts
func (_m *MockRobotClientService) GetRobotByID(params *robot.GetRobotByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...robot.ClientOption) (*robot.GetRobotByIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *robot.GetRobotByIDOK
	if rf, ok := ret.Get(0).(func(*robot.GetRobotByIDParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) *robot.GetRobotByIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.GetRobotByIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.GetRobotByIDParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRobot provides a mock function with given fields: params, authInfo, opts
func (_m *MockRobotClientService) ListRobot(params *robot.ListRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...robot.ClientOption) (*robot.ListRobotOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *robot.ListRobotOK
	if rf, ok := ret.Get(0).(func(*robot.ListRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) *robot.ListRobotOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.ListRobotOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.ListRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshSec provides a mock function with given fields: params, authInfo, opts
func (_m *MockRobotClientService) RefreshSec(params *robot.RefreshSecParams, authInfo runtime.ClientAuthInfoWriter, opts ...robot.ClientOption) (*robot.RefreshSecOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *robot.RefreshSecOK
	if rf, ok := ret.Get(0).(func(*robot.RefreshSecParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) *robot.RefreshSecOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.RefreshSecOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.RefreshSecParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockRobotClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateRobot provides a mock function with given fields: params, authInfo, opts
func (_m *MockRobotClientService) UpdateRobot(params *robot.UpdateRobotParams, authInfo runtime.ClientAuthInfoWriter, opts ...robot.ClientOption) (*robot.UpdateRobotOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *robot.UpdateRobotOK
	if rf, ok := ret.Get(0).(func(*robot.UpdateRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) *robot.UpdateRobotOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*robot.UpdateRobotOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*robot.UpdateRobotParams, runtime.ClientAuthInfoWriter, ...robot.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
