package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	models "github.com/lp/music-chart/internal/model"
)

// GetToken function
func GetToken(data *models.User) (token string, err error) {
	exp, _ := strconv.ParseInt(os.Getenv("AUTH_EXPIRATION"), 10, 64)
	expirationDuration := time.Duration(exp) * time.Hour

	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationDuration).Unix(),
		},
		UserID: data.ID,
	}

	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	token, err = sign.SignedString([]byte(os.Getenv("AUTH_KEY")))
	if err != nil {
		return
	}

	return
}
