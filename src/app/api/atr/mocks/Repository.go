// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import helpers "github.com/atomic/sip/src/helpers"
import mock "github.com/stretchr/testify/mock"
import models "github.com/atomic/sip/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *Repository) Create(_a0 models.Brand) (*models.Brand, error) {
	ret := _m.Called(_a0)

	var r0 *models.Brand
	if rf, ok := ret.Get(0).(func(models.Brand) *models.Brand); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Brand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Brand) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0
func (_m *Repository) Delete(_a0 string) (*models.Brand, error) {
	ret := _m.Called(_a0)

	var r0 *models.Brand
	if rf, ok := ret.Get(0).(func(string) *models.Brand); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Brand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: _a0
func (_m *Repository) Find(_a0 string) (*models.Brand, error) {
	ret := _m.Called(_a0)

	var r0 *models.Brand
	if rf, ok := ret.Get(0).(func(string) *models.Brand); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Brand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: _a0
func (_m *Repository) FindAll(_a0 helpers.FindAllParams) ([]*models.Brand, int, error) {
	ret := _m.Called(_a0)

	var r0 []*models.Brand
	if rf, ok := ret.Get(0).(func(helpers.FindAllParams) []*models.Brand); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Brand)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(helpers.FindAllParams) int); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(helpers.FindAllParams) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindStatus provides a mock function with given fields:
func (_m *Repository) FindStatus() ([]models.BrandStatus, error) {
	ret := _m.Called()

	var r0 []models.BrandStatus
	if rf, ok := ret.Get(0).(func() []models.BrandStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BrandStatus)
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

// Update provides a mock function with given fields: _a0, _a1
func (_m *Repository) Update(_a0 string, _a1 models.Brand) (*models.Brand, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Brand
	if rf, ok := ret.Get(0).(func(string, models.Brand) *models.Brand); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Brand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, models.Brand) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
