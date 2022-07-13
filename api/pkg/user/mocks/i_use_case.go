// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "shared-bike/domain"

	mock "github.com/stretchr/testify/mock"
)

// IUseCase is an autogenerated mock type for the IUseCase type
type IUseCase struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, body
func (_m *IUseCase) Login(ctx context.Context, body domain.LoginBody) (domain.UserDTO, error) {
	ret := _m.Called(ctx, body)

	var r0 domain.UserDTO
	if rf, ok := ret.Get(0).(func(context.Context, domain.LoginBody) domain.UserDTO); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Get(0).(domain.UserDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.LoginBody) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, body
func (_m *IUseCase) Register(ctx context.Context, body domain.RegisterBody) (domain.UserDTO, error) {
	ret := _m.Called(ctx, body)

	var r0 domain.UserDTO
	if rf, ok := ret.Get(0).(func(context.Context, domain.RegisterBody) domain.UserDTO); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Get(0).(domain.UserDTO)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RegisterBody) error); ok {
		r1 = rf(ctx, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUseCase creates a new instance of IUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUseCase(t mockConstructorTestingTNewIUseCase) *IUseCase {
	mock := &IUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}