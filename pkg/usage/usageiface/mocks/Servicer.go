// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import usage "github.com/Optum/dce/pkg/usage"

// Servicer is an autogenerated mock type for the Servicer type
type Servicer struct {
	mock.Mock
}

// GetLease provides a mock function with given fields: id
func (_m *Servicer) GetLease(id string) (*usage.Lease, error) {
	ret := _m.Called(id)

	var r0 *usage.Lease
	if rf, ok := ret.Get(0).(func(string) *usage.Lease); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usage.Lease)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLease provides a mock function with given fields: query
func (_m *Servicer) ListLease(query *usage.Lease) (*usage.Leases, error) {
	ret := _m.Called(query)

	var r0 *usage.Leases
	if rf, ok := ret.Get(0).(func(*usage.Lease) *usage.Leases); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usage.Leases)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usage.Lease) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPrincipal provides a mock function with given fields: query
func (_m *Servicer) ListPrincipal(query *usage.Principal) (*usage.Principal, error) {
	ret := _m.Called(query)

	var r0 *usage.Principal
	if rf, ok := ret.Get(0).(func(*usage.Principal) *usage.Principal); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usage.Principal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usage.Principal) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpsertLeaseUsage provides a mock function with given fields: data
func (_m *Servicer) UpsertLeaseUsage(data *usage.Lease) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(*usage.Lease) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}