package usecase

import (
	"context"
	"errors"

	models "github.com/lp/music-chart/internal/model"
)

var handleGetUserInfo = getUserInfo

func getUserInfo(ctx context.Context) (*models.UserInfo, error) {
	userInfo, isValid := ctx.Value("userInfo").(models.UserInfo)
	if !isValid {
		err := errors.New("struct not valid")
		return nil, err
	}

	return &userInfo, nil
}
