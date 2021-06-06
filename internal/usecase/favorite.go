package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	models "github.com/lp/music-chart/internal/model"
	r "github.com/lp/music-chart/internal/repo/favorite"
	mR "github.com/lp/music-chart/internal/repo/music"
)

// FavoriteUsecase interface
type FavoriteUsecase interface {
	SetFavorite(ctx context.Context, id uint64) (*models.SetFavoriteResponse, int, error)
}

// FavoriteUsecaseImpl struct
type FavoriteUsecaseImpl struct {
	FavoriteRepo r.Repository
	MusicRepo    mR.Repository
}

// NewFavoriteUsecaseImpl function
func NewFavoriteUsecaseImpl(params FavoriteUsecaseImpl) FavoriteUsecase {
	return &params
}

// SetFavorite function
func (f *FavoriteUsecaseImpl) SetFavorite(ctx context.Context, id uint64) (*models.SetFavoriteResponse, int, error) {
	userInfo, isValid := ctx.Value("userInfo").(models.UserInfo)
	if !isValid {
		err := errors.New("struct not valid")
		return nil, http.StatusInternalServerError, err
	}

	_, err := f.MusicRepo.GetByID(ctx, id)
	if err == sql.ErrNoRows {
		err = errors.New("music is not found")
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	affectRow, err := f.FavoriteRepo.Delete(ctx, userInfo.ID, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if affectRow == 0 {
		err = f.FavoriteRepo.Insert(ctx, models.Favorite{
			UserID:  userInfo.ID,
			MusicID: id,
		})
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	return &models.SetFavoriteResponse{
		IsFavorite: affectRow == 0,
	}, http.StatusOK, nil
}
