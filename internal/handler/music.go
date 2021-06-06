package handler

import (
	"net/http"
	"strconv"
	"time"

	m "github.com/lp/music-chart/internal/model"

	"github.com/lp/music-chart/util"

	"github.com/labstack/echo"
)

func (h *handler) GetMusicChart(c echo.Context) (err error) {
	start := time.Now()
	resp := m.Response{
		Status: http.StatusOK,
	}

	// Bind Request
	qParam := util.BindQueryParams(c.Request())

	// Sent Data to UseCase and get process result
	res, err := h.ucs.UcMusic.GetAll(util.EchoToContext(c, "userInfo"), qParam)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusInternalServerError
		resp.ProcessTime = util.Float64ToString(time.Since(start).Seconds())
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp = m.Response{
		Data:        res,
		ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
		Status:      http.StatusOK,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetMusicChartDetail(c echo.Context) (err error) {
	start := time.Now()
	resp := m.Response{
		Status: http.StatusOK,
	}

	// Bind Request
	id := c.Param("id")
	uID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = http.StatusBadRequest
		resp.ProcessTime = util.Float64ToString(time.Since(start).Seconds())
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// Sent Data to UseCase and get process result
	res, httpCode, err := h.ucs.UcMusic.GetByID(util.EchoToContext(c, "userInfo"), uID)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = httpCode
		resp.ProcessTime = util.Float64ToString(time.Since(start).Seconds())
		return c.JSON(httpCode, resp)
	}

	resp = m.Response{
		Data:        res,
		ProcessTime: util.Float64ToString(time.Since(start).Seconds()),
		Status:      http.StatusOK,
	}

	return c.JSON(http.StatusOK, resp)
}
