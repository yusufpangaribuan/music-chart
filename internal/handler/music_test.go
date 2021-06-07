package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	m "github.com/lp/music-chart/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAPI_GetMusicChart(t *testing.T) {
	setup()
	mockErr := errors.New("error mock")
	path := "/music-chart"

	tests := []struct {
		name  string
		patch func() echo.Context
		want  int
	}{
		{
			name: "When_GetAllReturnError_Expect_StatusInternalServerError",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)

				handleBindQueryParams = func(req *http.Request) (params m.BasicSelectParams) {
					return
				}

				handleEchoToContext = func(c echo.Context, key string) context.Context {
					return context.TODO()
				}

				mockMusicUC.On("GetAll", context.TODO(), m.BasicSelectParams{}).Return(nil, mockErr).Once()
				return ctx
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "When_GetMusicChartReturnSuccess_Expect_StatusOK",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)

				handleBindQueryParams = func(req *http.Request) (params m.BasicSelectParams) {
					return
				}

				handleEchoToContext = func(c echo.Context, key string) context.Context {
					return context.TODO()
				}

				mockMusicUC.On("GetAll", context.TODO(), m.BasicSelectParams{}).Return(nil, nil).Once()
				return ctx
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHandler(
				Usecases{
					UcMusic: mockMusicUC,
				},
			)
			ctx := tt.patch()
			_ = u.GetMusicChart(ctx)
			assert.Equal(t, tt.want, ctx.Response().Status)
		})
	}
}

func TestAPI_GetMusicChartDetail(t *testing.T) {
	setup()
	mockErr := errors.New("error mock")
	path := "/music-chart"

	tests := []struct {
		name  string
		patch func() echo.Context
		want  int
	}{
		{
			name: "When_ParseUintReturnError_Expect_StatusBadRequest",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)
				ctx.SetParamNames("id")
				ctx.SetParamValues("error")
				return ctx
			},
			want: http.StatusBadRequest,
		},
		{
			name: "When_GetByIDReturnError_Expect_StatusInternalServerError",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				handleEchoToContext = func(c echo.Context, key string) context.Context {
					return context.TODO()
				}

				mockMusicUC.On("GetByID", context.TODO(), uint64(1)).Return(nil, http.StatusInternalServerError, mockErr).Once()

				return ctx
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "When_GetMusicChartDetailReturnSuccess_Expect_StatusOK",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				handleEchoToContext = func(c echo.Context, key string) context.Context {
					return context.TODO()
				}

				mockMusicUC.On("GetByID", context.TODO(), uint64(1)).Return(nil, http.StatusOK, nil).Once()

				return ctx
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHandler(
				Usecases{
					UcMusic: mockMusicUC,
				},
			)
			ctx := tt.patch()
			_ = u.GetMusicChartDetail(ctx)
			assert.Equal(t, tt.want, ctx.Response().Status)
		})
	}
}
