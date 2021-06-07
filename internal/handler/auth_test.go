package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	m "github.com/lp/music-chart/internal/model"
	"github.com/lp/music-chart/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

var mockAuthUC *mocks.Auth
var mockFavoriteUC *mocks.FavoriteUsecase
var mockMusicUC *mocks.MusicUsecase

func setup() {
	mockAuthUC = new(mocks.Auth)
	mockFavoriteUC = new(mocks.FavoriteUsecase)
	mockMusicUC = new(mocks.MusicUsecase)
}

func TestAPI_Login(t *testing.T) {
	setup()
	mockErr := errors.New("error mock")
	path := "/login"

	tests := []struct {
		name  string
		patch func() echo.Context
		want  int
	}{
		{
			name: "When_DecodeReturnError_Expect_StatusBadRequest",
			patch: func() echo.Context {
				req := `[]`
				ctx := getMockEchoContext(req, http.MethodPost, path)
				return ctx
			},
			want: http.StatusBadRequest,
		},
		{
			name: "When_AuthLoginReturnError_Expect_StatusBadRequest",
			patch: func() echo.Context {
				req := m.LoginReq{
					Password: "test",
					UserName: "test",
				}
				b, _ := json.Marshal(req)
				ctx := getMockEchoContext(string(b), http.MethodPost, path)

				mockAuthUC.On("Login", context.TODO(), req).Return("", http.StatusBadRequest, mockErr).Once()

				return ctx
			},
			want: http.StatusBadRequest,
		},
		{
			name: "When_AuthLoginReturnSuccess_Expect_StatusOK",
			patch: func() echo.Context {
				req := m.LoginReq{
					Password: "test",
					UserName: "test",
				}
				b, _ := json.Marshal(req)
				ctx := getMockEchoContext(string(b), http.MethodPost, path)

				mockAuthUC.On("Login", context.TODO(), req).Return("test", http.StatusOK, nil).Once()

				return ctx
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHandler(
				Usecases{
					UcAuth: mockAuthUC,
				},
			)
			ctx := tt.patch()
			_ = u.Login(ctx)
			assert.Equal(t, tt.want, ctx.Response().Status)
		})
	}
}

func TestAPI_Register(t *testing.T) {
	setup()
	mockErr := errors.New("error mock")
	path := "/register"

	tests := []struct {
		name  string
		patch func() echo.Context
		want  int
	}{
		{
			name: "When_DecodeReturnError_Expect_StatusBadRequest",
			patch: func() echo.Context {
				req := `[]`
				ctx := getMockEchoContext(req, http.MethodPost, path)
				return ctx
			},
			want: http.StatusBadRequest,
		},
		{
			name: "When_handleRegisterValidationReturnError_Expect_StatusBadRequest",
			patch: func() echo.Context {
				req := m.User{
					Password: "test",
					UserName: "test",
				}
				b, _ := json.Marshal(req)
				ctx := getMockEchoContext(string(b), http.MethodPost, path)
				handleRegisterValidation = func(regReq m.User) (err error) {
					return mockErr
				}
				return ctx
			},
			want: http.StatusBadRequest,
		},
		{
			name: "When_RegisterReturnError_Expect_StatusInternalServerError",
			patch: func() echo.Context {
				req := m.User{
					Password: "test",
					UserName: "test",
				}
				b, _ := json.Marshal(req)
				ctx := getMockEchoContext(string(b), http.MethodPost, path)
				handleRegisterValidation = func(regReq m.User) (err error) {
					return nil
				}
				mockAuthUC.On("Register", context.TODO(), req).Return("", http.StatusInternalServerError, mockErr).Once()

				return ctx
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "When_AuthRegisterReturnSuccess_Expect_StatusOK",
			patch: func() echo.Context {
				req := m.User{
					Password: "test",
					UserName: "test",
				}
				b, _ := json.Marshal(req)
				ctx := getMockEchoContext(string(b), http.MethodPost, path)
				handleRegisterValidation = func(regReq m.User) (err error) {
					return nil
				}
				mockAuthUC.On("Register", context.TODO(), req).Return("test", http.StatusOK, nil).Once()

				return ctx
			},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewHandler(
				Usecases{
					UcAuth: mockAuthUC,
				},
			)
			ctx := tt.patch()
			_ = u.Register(ctx)
			assert.Equal(t, tt.want, ctx.Response().Status)
		})
	}
}

func getMockEchoContext(userJSON, method, path string) echo.Context {

	e := echo.New()
	req := httptest.NewRequest(method, path, strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c
}
