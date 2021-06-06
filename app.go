package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	conMysql "github.com/lp/music-chart/connections/mysql"
	"github.com/lp/music-chart/internal/handler"
	"github.com/lp/music-chart/internal/middleware"
)

var (
	repos        repo
	usecases     usecase
	fDevelopment bool
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {

	// Init connection
	conMysql.Initialize()

	// Init all depedencies for repo
	initRepo()
	initUsecase()

	// Init handler / controller
	handler := handler.NewHandler(handler.Usecases{
		UcAuth:     usecases.user,
		UcFavorite: usecases.favorite,
		UcMusic:    usecases.music,
	})

	e := echo.New()
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		Skipper:      echoMiddleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	withoutToken := e.Group("/lp")
	// Login & register
	withoutToken.POST("/login", handler.Login)
	withoutToken.POST("/register", handler.Register)
	authModule := middleware.New()
	r := e.Group("/lp")
	r.Use(authModule.Auth)
	// Core Function
	r.GET("/music-chart", handler.GetMusicChart)
	r.GET("/music-chart/:id", handler.GetMusicChartDetail)
	r.POST("/music-chart/favorite/:id", handler.Favorite)

	e.Logger.Fatal(e.Start(os.Getenv("DEFAULT_PORT")))
}
