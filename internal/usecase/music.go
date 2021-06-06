package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	models "github.com/lp/music-chart/internal/model"
	r "github.com/lp/music-chart/internal/repo/music"
)

// MusicUsecase interface
type MusicUsecase interface {
	GetByID(ctx context.Context, id uint64) (res *models.Music, httpCode int, err error)
	GetAll(ctx context.Context, params models.BasicSelectParams) (res *models.GetAllResponse, err error)
}

// MusicUsecaseImpl struct
type MusicUsecaseImpl struct {
	MusicRepo r.Repository
}

// NewMusicUsecaseImpl function
func NewMusicUsecaseImpl(params MusicUsecaseImpl) MusicUsecase {
	return &params
}

// GetByID this function call music repo to get music data by music id
func (m *MusicUsecaseImpl) GetByID(ctx context.Context, id uint64) (res *models.Music, httpCode int, err error) {
	httpCode = http.StatusOK

	userInfo, isValid := ctx.Value("userInfo").(models.UserInfo)
	if !isValid {
		err = errors.New("struct not valid")
		httpCode = http.StatusBadRequest
		return
	}

	data, err := m.MusicRepo.GetByID(ctx, id)
	if err == sql.ErrNoRows {
		return nil, http.StatusNotFound, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Set is_favorite value
	data.IsFavorite = data.FavoriteID.Uint64 == userInfo.ID

	return data, httpCode, nil
}

// GetAll this function call music repo to get all music data
func (m *MusicUsecaseImpl) GetAll(ctx context.Context, params models.BasicSelectParams) (res *models.GetAllResponse, err error) {
	userInfo, isValid := ctx.Value("userInfo").(models.UserInfo)
	if !isValid {
		err = errors.New("struct not valid")
		return
	}

	data, tot, err := m.MusicRepo.GetAll(ctx, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}

	// Set is_favorite value
	m.setIsFavorite(data, userInfo.ID)

	return &models.GetAllResponse{
		Items: data,
		Page:  params.Page,
		Total: tot,
	}, nil
}

func (m *MusicUsecaseImpl) setIsFavorite(data []*models.Music, userID uint64) {
	for _, val := range data {
		val.IsFavorite = val.FavoriteID.Uint64 == userID
	}
}
