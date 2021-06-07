// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/lp/music-chart/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, req
func (_m *Auth) Login(ctx context.Context, req models.LoginReq) (string, int, error) {
	ret := _m.Called(ctx, req)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, models.LoginReq) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, models.LoginReq) int); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, models.LoginReq) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: ctx, req
func (_m *Auth) Register(ctx context.Context, req models.User) (string, int, error) {
	ret := _m.Called(ctx, req)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, models.User) string); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, models.User) int); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, models.User) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}