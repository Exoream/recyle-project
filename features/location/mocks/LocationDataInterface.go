// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	entity "recycle/features/location/entity"

	mock "github.com/stretchr/testify/mock"
)

// LocationDataInterface is an autogenerated mock type for the LocationDataInterface type
type LocationDataInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *LocationDataInterface) Create(data entity.Main) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Main) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: id
func (_m *LocationDataInterface) DeleteById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllLocation provides a mock function with given fields:
func (_m *LocationDataInterface) FindAllLocation() ([]entity.Main, error) {
	ret := _m.Called()

	var r0 []entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Main, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Main); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Main)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCity provides a mock function with given fields: city
func (_m *LocationDataInterface) GetByCity(city string) ([]entity.Main, error) {
	ret := _m.Called(city)

	var r0 []entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]entity.Main, error)); ok {
		return rf(city)
	}
	if rf, ok := ret.Get(0).(func(string) []entity.Main); ok {
		r0 = rf(city)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Main)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *LocationDataInterface) GetById(id string) (entity.Main, error) {
	ret := _m.Called(id)

	var r0 entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Main, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Main); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.Main)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: id, updated
func (_m *LocationDataInterface) UpdateById(id string, updated entity.Main) (entity.Main, error) {
	ret := _m.Called(id, updated)

	var r0 entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func(string, entity.Main) (entity.Main, error)); ok {
		return rf(id, updated)
	}
	if rf, ok := ret.Get(0).(func(string, entity.Main) entity.Main); ok {
		r0 = rf(id, updated)
	} else {
		r0 = ret.Get(0).(entity.Main)
	}

	if rf, ok := ret.Get(1).(func(string, entity.Main) error); ok {
		r1 = rf(id, updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLocationDataInterface creates a new instance of LocationDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLocationDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *LocationDataInterface {
	mock := &LocationDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}