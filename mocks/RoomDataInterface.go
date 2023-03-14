// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	rooms "groupproject3-airbnb-api/features/rooms"

	mock "github.com/stretchr/testify/mock"
)

// RoomDataInterface is an autogenerated mock type for the RoomDataInterface type
type RoomDataInterface struct {
	mock.Mock
}

// Destroy provides a mock function with given fields: id
func (_m *RoomDataInterface) Destroy(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Edit provides a mock function with given fields: roomEntity, id
func (_m *RoomDataInterface) Edit(roomEntity rooms.RoomEntity, id uint) error {
	ret := _m.Called(roomEntity, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(rooms.RoomEntity, uint) error); ok {
		r0 = rf(roomEntity, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields:
func (_m *RoomDataInterface) SelectAll() ([]rooms.RoomEntity, error) {
	ret := _m.Called()

	var r0 []rooms.RoomEntity
	if rf, ok := ret.Get(0).(func() []rooms.RoomEntity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rooms.RoomEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *RoomDataInterface) SelectById(id uint) (rooms.RoomEntity, error) {
	ret := _m.Called(id)

	var r0 rooms.RoomEntity
	if rf, ok := ret.Get(0).(func(uint) rooms.RoomEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(rooms.RoomEntity)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByUserId provides a mock function with given fields: user_id
func (_m *RoomDataInterface) SelectByUserId(user_id uint) ([]rooms.RoomEntity, error) {
	ret := _m.Called(user_id)

	var r0 []rooms.RoomEntity
	if rf, ok := ret.Get(0).(func(uint) []rooms.RoomEntity); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]rooms.RoomEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: roomEntity
func (_m *RoomDataInterface) Store(roomEntity rooms.RoomEntity) (uint, error) {
	ret := _m.Called(roomEntity)

	var r0 uint
	if rf, ok := ret.Get(0).(func(rooms.RoomEntity) uint); ok {
		r0 = rf(roomEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(rooms.RoomEntity) error); ok {
		r1 = rf(roomEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRoomDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoomDataInterface creates a new instance of RoomDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoomDataInterface(t mockConstructorTestingTNewRoomDataInterface) *RoomDataInterface {
	mock := &RoomDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
