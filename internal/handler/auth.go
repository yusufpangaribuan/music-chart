package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
	m "github.com/lp/music-chart/internal/model"
	"github.com/lp/music-chart/util"
)

func (h *handler) Login(c echo.Context) (err error) {
	start := time.Now()

	// Bind Request Body
	var loginreq m.LoginReq
	err = json.NewDecoder(c.Request().Body).Decode(&loginreq)
	if err != nil {
		resp := m.Response{
			ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
			Error:       err.Error(),
			Status:      http.StatusBadRequest,
		}

		return c.JSON(http.StatusBadRequest, resp)
	}

	// Sent Data to UseCase and get process result
	token, httpCode, err := h.ucs.UcAuth.Login(context.TODO(), loginreq)
	if err != nil {
		return c.JSON(httpCode, m.Response{
			ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
			Error:       err.Error(),
			Status:      httpCode,
		})
	}

	return c.JSON(http.StatusOK, m.Response{
		Data:        token,
		ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
		Status:      http.StatusOK,
	})
}

func (h *handler) Register(c echo.Context) (err error) {
	start := time.Now()

	// Bind Request Body
	var regReq m.User
	err = json.NewDecoder(c.Request().Body).Decode(&regReq)
	if err != nil {
		resp := m.Response{
			ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
			Error:       err.Error(),
			Status:      http.StatusBadRequest,
		}

		return c.JSON(http.StatusBadRequest, resp)
	}

	// Request Validation
	err = handleRegisterValidation(regReq)
	if err != nil {
		resp := m.Response{
			ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
			Error:       err.Error(),
			Status:      http.StatusBadRequest,
		}

		return c.JSON(http.StatusBadRequest, resp)
	}

	// Sent Data to UseCase and get process result
	token, httpCode, err := h.ucs.UcAuth.Register(context.TODO(), regReq)
	if err != nil {
		return c.JSON(httpCode, m.Response{
			ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
			Error:       err.Error(),
			Status:      httpCode,
		})
	}

	return c.JSON(http.StatusOK, m.Response{
		Data:        token,
		ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
		Status:      http.StatusOK,
	})
}
