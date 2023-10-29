// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	entity "recycle/features/pickup/entity"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// PickupDataInterface is an autogenerated mock type for the PickupDataInterface type
type PickupDataInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: data, image
func (_m *PickupDataInterface) Create(data entity.Main, image *multipart.FileHeader) error {
	ret := _m.Called(data, image)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Main, *multipart.FileHeader) error); ok {
		r0 = rf(data, image)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: id
func (_m *PickupDataInterface) DeleteById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllPickup provides a mock function with given fields:
func (_m *PickupDataInterface) FindAllPickup() ([]entity.Main, error) {
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

// GetById provides a mock function with given fields: id
func (_m *PickupDataInterface) GetById(id string) (entity.Main, error) {
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

// GetByStatus provides a mock function with given fields: status
func (_m *PickupDataInterface) GetByStatus(status string) ([]entity.Main, error) {
	ret := _m.Called(status)

	var r0 []entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]entity.Main, error)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) []entity.Main); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Main)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: id, updated, image
func (_m *PickupDataInterface) UpdateById(id string, updated entity.Main, image *multipart.FileHeader) (entity.Main, error) {
	ret := _m.Called(id, updated, image)

	var r0 entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func(string, entity.Main, *multipart.FileHeader) (entity.Main, error)); ok {
		return rf(id, updated, image)
	}
	if rf, ok := ret.Get(0).(func(string, entity.Main, *multipart.FileHeader) entity.Main); ok {
		r0 = rf(id, updated, image)
	} else {
		r0 = ret.Get(0).(entity.Main)
	}

	if rf, ok := ret.Get(1).(func(string, entity.Main, *multipart.FileHeader) error); ok {
		r1 = rf(id, updated, image)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: pickupID, newStatus
func (_m *PickupDataInterface) UpdateStatus(pickupID string, newStatus string) error {
	ret := _m.Called(pickupID, newStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(pickupID, newStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPickupDataInterface creates a new instance of PickupDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPickupDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PickupDataInterface {
	mock := &PickupDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
