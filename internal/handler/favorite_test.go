package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestAPI_Favorite(t *testing.T) {
	setup()
	mockErr := errors.New("error mock")
	path := "/music-chart/favorite/:id"

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
			name: "When_SetFavoriteReturnError_Expect_StatusInternalServerError",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				handleEchoToContext = func(c echo.Context, key string) context.Context {
					return context.TODO()
				}
				mockFavoriteUC.On("SetFavorite", context.TODO(), uint64(1)).Return(nil, http.StatusInternalServerError, mockErr).Once()

				return ctx
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "When_FavoriteReturnSuccess_Expect_StatusOK",
			patch: func() echo.Context {
				ctx := getMockEchoContext("", http.MethodGet, path)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")
				handleEchoToContext = func(c echo.Context, key string) context.Context {
					return context.TODO()
				}
				mockFavoriteUC.On("SetFavorite", context.TODO(), uint64(1)).Return(nil, http.StatusOK, nil).Once()

				return ctx
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHandler(
				Usecases{
					UcFavorite: mockFavoriteUC,
				},
			)
			ctx := tt.patch()
			_ = u.Favorite(ctx)
			assert.Equal(t, tt.want, ctx.Response().Status)
		})
	}
}
