package handler

import "github.com/lp/music-chart/internal/usecase"

// Usecases struct
type Usecases struct {
	UcAuth     usecase.Auth
	UcMusic    usecase.MusicUsecase
	UcFavorite usecase.FavoriteUsecase
}

type handler struct {
	ucs Usecases
}

func NewHandler(ucs Usecases) Handler {
	return &handler{ucs: ucs}
}
