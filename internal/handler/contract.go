package handler

import "github.com/labstack/echo"

type Handler interface {
	// Auth Related
	Login(c echo.Context) (err error)
	Register(c echo.Context) (err error)

	// Music Related
	GetMusicChart(c echo.Context) (err error)
	GetMusicChartDetail(c echo.Context) (err error)

	// Favorite Related
	Favorite(c echo.Context) (err error)
}
