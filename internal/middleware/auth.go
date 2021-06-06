package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	models "github.com/lp/music-chart/internal/model"
)

// Module struct
type Module struct {
}

// New function
func New() *Module {
	return &Module{}
}

// Auth function
func (m *Module) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			result := gin.H{
				"message": "Invalid token",
			}
			c.JSON(http.StatusBadRequest, result)
			return nil
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("AUTH_KEY")), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			result := gin.H{
				"message": "not authorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			return nil
		}

		c.Set("userInfo", models.UserInfo{
			ID: uint64(claims["UserID"].(float64)),
		})

		return next(c)
	}
}
