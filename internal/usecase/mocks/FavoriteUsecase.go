// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/lp/music-chart/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// FavoriteUsecase is an autogenerated mock type for the FavoriteUsecase type
type FavoriteUsecase struct {
	mock.Mock
}

// SetFavorite provides a mock function with given fields: ctx, id
func (_m *FavoriteUsecase) SetFavorite(ctx context.Context, id uint64) (*models.SetFavoriteResponse, int, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.SetFavoriteResponse
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.SetFavoriteResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SetFavoriteResponse)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, uint64) int); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, uint64) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
