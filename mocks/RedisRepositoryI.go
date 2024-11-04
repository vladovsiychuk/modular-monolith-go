// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	backendforfrontend "github.com/vladovsiychuk/microservice-demo-go/internal/backendforfrontend"

	uuid "github.com/google/uuid"
)

// RedisRepositoryI is an autogenerated mock type for the RedisRepositoryI type
type RedisRepositoryI struct {
	mock.Mock
}

// FindByPostId provides a mock function with given fields: _a0
func (_m *RedisRepositoryI) FindByPostId(_a0 uuid.UUID) (backendforfrontend.PostAggregateI, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FindByPostId")
	}

	var r0 backendforfrontend.PostAggregateI
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (backendforfrontend.PostAggregateI, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) backendforfrontend.PostAggregateI); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(backendforfrontend.PostAggregateI)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCache provides a mock function with given fields: _a0
func (_m *RedisRepositoryI) UpdateCache(_a0 backendforfrontend.PostAggregateI) {
	_m.Called(_a0)
}

// NewRedisRepositoryI creates a new instance of RedisRepositoryI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRedisRepositoryI(t interface {
	mock.TestingT
	Cleanup(func())
}) *RedisRepositoryI {
	mock := &RedisRepositoryI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
