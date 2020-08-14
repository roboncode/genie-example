// Code generated by mockery v1.1.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Get provides a mock function with given fields: name
func (_m *Store) Get(name string) interface{} {
	ret := _m.Called(name)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Set provides a mock function with given fields: name, value
func (_m *Store) Set(name string, value interface{}) {
	_m.Called(name, value)
}