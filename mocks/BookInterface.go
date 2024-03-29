// Code generated by mockery v2.39.2. DO NOT EDIT.

package mocks

import (
	entity "library/entity"

	mock "github.com/stretchr/testify/mock"
)

// BookInterface is an autogenerated mock type for the BookInterface type
type BookInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: book
func (_m *BookInterface) Create(book entity.Book) (entity.Book, error) {
	ret := _m.Called(book)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entity.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Book) (entity.Book, error)); ok {
		return rf(book)
	}
	if rf, ok := ret.Get(0).(func(entity.Book) entity.Book); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Get(0).(entity.Book)
	}

	if rf, ok := ret.Get(1).(func(entity.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: bookId
func (_m *BookInterface) Delete(bookId int) error {
	ret := _m.Called(bookId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(bookId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: book
func (_m *BookInterface) Get(book entity.Book) (entity.Book, error) {
	ret := _m.Called(book)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 entity.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Book) (entity.Book, error)); ok {
		return rf(book)
	}
	if rf, ok := ret.Get(0).(func(entity.Book) entity.Book); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Get(0).(entity.Book)
	}

	if rf, ok := ret.Get(1).(func(entity.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *BookInterface) List() ([]entity.Book, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []entity.Book
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Book, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Book)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: book
func (_m *BookInterface) Update(book entity.Book) (entity.Book, error) {
	ret := _m.Called(book)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 entity.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Book) (entity.Book, error)); ok {
		return rf(book)
	}
	if rf, ok := ret.Get(0).(func(entity.Book) entity.Book); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Get(0).(entity.Book)
	}

	if rf, ok := ret.Get(1).(func(entity.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBookInterface creates a new instance of BookInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookInterface {
	mock := &BookInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
