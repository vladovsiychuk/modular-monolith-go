// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	auth "github.com/vladovsiychuk/microservice-demo-go/internal/auth"
)

// KeyRepositoryI is an autogenerated mock type for the KeyRepositoryI type
type KeyRepositoryI struct {
	mock.Mock
}

// GetKeys provides a mock function with given fields:
func (_m *KeyRepositoryI) GetKeys() (auth.KeysI, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetKeys")
	}

	var r0 auth.KeysI
	var r1 error
	if rf, ok := ret.Get(0).(func() (auth.KeysI, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() auth.KeysI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(auth.KeysI)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *KeyRepositoryI) Update(_a0 auth.KeysI) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(auth.KeysI) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKeyRepositoryI creates a new instance of KeyRepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKeyRepositoryI(t interface {
	mock.TestingT
	Cleanup(func())
}) *KeyRepositoryI {
	mock := &KeyRepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
