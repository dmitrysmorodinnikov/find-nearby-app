// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	model "beam-backend/model"

	mock "github.com/stretchr/testify/mock"
)

// LocationRepository is an autogenerated mock type for the LocationRepository type
type LocationRepository struct {
	mock.Mock
}

// FindVehicleLocations provides a mock function with given fields: latitude, longitude, radius, limit
func (_m *LocationRepository) FindVehicleLocations(latitude float64, longitude float64, radius int, limit int) ([]model.Location, error) {
	ret := _m.Called(latitude, longitude, radius, limit)

	var r0 []model.Location
	if rf, ok := ret.Get(0).(func(float64, float64, int, int) []model.Location); ok {
		r0 = rf(latitude, longitude, radius, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float64, float64, int, int) error); ok {
		r1 = rf(latitude, longitude, radius, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
