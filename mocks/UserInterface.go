// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	entity "library/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserInterface is an autogenerated mock type for the UserInterface type
type UserInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserInterface) Create(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (entity.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userId
func (_m *UserInterface) Delete(userId int) error {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: username
func (_m *UserInterface) Get(username string) (entity.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) entity.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *UserInterface) List() ([]entity.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: user
func (_m *UserInterface) Update(user entity.User) (entity.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.User) (entity.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(entity.User) entity.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserInterface creates a new instance of UserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserInterface {
	mock := &UserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
