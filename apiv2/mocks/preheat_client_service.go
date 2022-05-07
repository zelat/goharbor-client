// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	preheat "github.com/zelat/goharbor-client/v4/apiv2/internal/api/client/preheat"
	mock "github.com/stretchr/testify/mock"
)

// MockPreheatClientService is an autogenerated mock type for the ClientService type
type MockPreheatClientService struct {
	mock.Mock
}

// CreateInstance provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) CreateInstance(params *preheat.CreateInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.CreateInstanceCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.CreateInstanceCreated
	if rf, ok := ret.Get(0).(func(*preheat.CreateInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.CreateInstanceCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.CreateInstanceCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.CreateInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePolicy provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) CreatePolicy(params *preheat.CreatePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.CreatePolicyCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.CreatePolicyCreated
	if rf, ok := ret.Get(0).(func(*preheat.CreatePolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.CreatePolicyCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.CreatePolicyCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.CreatePolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteInstance provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) DeleteInstance(params *preheat.DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.DeleteInstanceOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.DeleteInstanceOK
	if rf, ok := ret.Get(0).(func(*preheat.DeleteInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.DeleteInstanceOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.DeleteInstanceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.DeleteInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePolicy provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) DeletePolicy(params *preheat.DeletePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.DeletePolicyOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.DeletePolicyOK
	if rf, ok := ret.Get(0).(func(*preheat.DeletePolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.DeletePolicyOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.DeletePolicyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.DeletePolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExecution provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) GetExecution(params *preheat.GetExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.GetExecutionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.GetExecutionOK
	if rf, ok := ret.Get(0).(func(*preheat.GetExecutionParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.GetExecutionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.GetExecutionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.GetExecutionParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInstance provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) GetInstance(params *preheat.GetInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.GetInstanceOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.GetInstanceOK
	if rf, ok := ret.Get(0).(func(*preheat.GetInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.GetInstanceOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.GetInstanceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.GetInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPolicy provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) GetPolicy(params *preheat.GetPolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.GetPolicyOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.GetPolicyOK
	if rf, ok := ret.Get(0).(func(*preheat.GetPolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.GetPolicyOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.GetPolicyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.GetPolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPreheatLog provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) GetPreheatLog(params *preheat.GetPreheatLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.GetPreheatLogOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.GetPreheatLogOK
	if rf, ok := ret.Get(0).(func(*preheat.GetPreheatLogParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.GetPreheatLogOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.GetPreheatLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.GetPreheatLogParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListExecutions provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ListExecutions(params *preheat.ListExecutionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ListExecutionsOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ListExecutionsOK
	if rf, ok := ret.Get(0).(func(*preheat.ListExecutionsParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ListExecutionsOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ListExecutionsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ListExecutionsParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListInstances provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ListInstances(params *preheat.ListInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ListInstancesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ListInstancesOK
	if rf, ok := ret.Get(0).(func(*preheat.ListInstancesParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ListInstancesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ListInstancesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ListInstancesParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPolicies provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ListPolicies(params *preheat.ListPoliciesParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ListPoliciesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ListPoliciesOK
	if rf, ok := ret.Get(0).(func(*preheat.ListPoliciesParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ListPoliciesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ListPoliciesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ListPoliciesParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviders provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ListProviders(params *preheat.ListProvidersParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ListProvidersOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ListProvidersOK
	if rf, ok := ret.Get(0).(func(*preheat.ListProvidersParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ListProvidersOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ListProvidersOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ListProvidersParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProvidersUnderProject provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ListProvidersUnderProject(params *preheat.ListProvidersUnderProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ListProvidersUnderProjectOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ListProvidersUnderProjectOK
	if rf, ok := ret.Get(0).(func(*preheat.ListProvidersUnderProjectParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ListProvidersUnderProjectOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ListProvidersUnderProjectOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ListProvidersUnderProjectParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTasks provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ListTasks(params *preheat.ListTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ListTasksOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ListTasksOK
	if rf, ok := ret.Get(0).(func(*preheat.ListTasksParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ListTasksOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ListTasksOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ListTasksParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManualPreheat provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) ManualPreheat(params *preheat.ManualPreheatParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.ManualPreheatCreated, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.ManualPreheatCreated
	if rf, ok := ret.Get(0).(func(*preheat.ManualPreheatParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.ManualPreheatCreated); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.ManualPreheatCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.ManualPreheatParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingInstances provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) PingInstances(params *preheat.PingInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.PingInstancesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.PingInstancesOK
	if rf, ok := ret.Get(0).(func(*preheat.PingInstancesParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.PingInstancesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.PingInstancesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.PingInstancesParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockPreheatClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// StopExecution provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) StopExecution(params *preheat.StopExecutionParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.StopExecutionOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.StopExecutionOK
	if rf, ok := ret.Get(0).(func(*preheat.StopExecutionParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.StopExecutionOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.StopExecutionOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.StopExecutionParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateInstance provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) UpdateInstance(params *preheat.UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.UpdateInstanceOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.UpdateInstanceOK
	if rf, ok := ret.Get(0).(func(*preheat.UpdateInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.UpdateInstanceOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.UpdateInstanceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.UpdateInstanceParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePolicy provides a mock function with given fields: params, authInfo, opts
func (_m *MockPreheatClientService) UpdatePolicy(params *preheat.UpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter, opts ...preheat.ClientOption) (*preheat.UpdatePolicyOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *preheat.UpdatePolicyOK
	if rf, ok := ret.Get(0).(func(*preheat.UpdatePolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) *preheat.UpdatePolicyOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*preheat.UpdatePolicyOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*preheat.UpdatePolicyParams, runtime.ClientAuthInfoWriter, ...preheat.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}