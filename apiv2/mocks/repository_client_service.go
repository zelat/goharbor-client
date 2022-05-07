// Code generated by mockery v2.7.5. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	repository "github.com/zelat/goharbor-client/v4/apiv2/internal/api/client/repository"
	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryClientService is an autogenerated mock type for the ClientService type
type MockRepositoryClientService struct {
	mock.Mock
}

// DeleteRepository provides a mock function with given fields: params, authInfo, opts
func (_m *MockRepositoryClientService) DeleteRepository(params *repository.DeleteRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...repository.ClientOption) (*repository.DeleteRepositoryOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *repository.DeleteRepositoryOK
	if rf, ok := ret.Get(0).(func(*repository.DeleteRepositoryParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) *repository.DeleteRepositoryOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.DeleteRepositoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.DeleteRepositoryParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepository provides a mock function with given fields: params, authInfo, opts
func (_m *MockRepositoryClientService) GetRepository(params *repository.GetRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...repository.ClientOption) (*repository.GetRepositoryOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *repository.GetRepositoryOK
	if rf, ok := ret.Get(0).(func(*repository.GetRepositoryParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) *repository.GetRepositoryOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.GetRepositoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.GetRepositoryParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositories provides a mock function with given fields: params, authInfo, opts
func (_m *MockRepositoryClientService) ListRepositories(params *repository.ListRepositoriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...repository.ClientOption) (*repository.ListRepositoriesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *repository.ListRepositoriesOK
	if rf, ok := ret.Get(0).(func(*repository.ListRepositoriesParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) *repository.ListRepositoriesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.ListRepositoriesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.ListRepositoriesParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockRepositoryClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateRepository provides a mock function with given fields: params, authInfo, opts
func (_m *MockRepositoryClientService) UpdateRepository(params *repository.UpdateRepositoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...repository.ClientOption) (*repository.UpdateRepositoryOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *repository.UpdateRepositoryOK
	if rf, ok := ret.Get(0).(func(*repository.UpdateRepositoryParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) *repository.UpdateRepositoryOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.UpdateRepositoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.UpdateRepositoryParams, runtime.ClientAuthInfoWriter, ...repository.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
