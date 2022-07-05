// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/duong-se/shared-bike/domain"
	mock "github.com/stretchr/testify/mock"
)

// IRepository is an autogenerated mock type for the IRepository type
type IRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *IRepository) GetByID(ctx context.Context, id int64) (*domain.Bike, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Bike
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.Bike); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Bike)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: ctx
func (_m *IRepository) GetList(ctx context.Context) (*[]domain.Bike, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.Bike
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.Bike); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Bike)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, payload
func (_m *IRepository) Update(ctx context.Context, payload *domain.Bike) error {
	ret := _m.Called(ctx, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Bike) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepository creates a new instance of IRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepository(t mockConstructorTestingTNewIRepository) *IRepository {
	mock := &IRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}