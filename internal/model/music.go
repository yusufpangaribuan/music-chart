package models

import (
	"github.com/volatiletech/null"
)

// Music struct
type Music struct {
	ID          uint64      `json:"id"`
	Title       string      `json:"title"`
	Singer      string      `json:"singer,omitempty"`
	Duration    string      `json:"duration,omitempty"`
	Album       string      `json:"album,omitempty"`
	ReleaseYear string      `json:"release_year,omitempty"`
	IsFavorite  bool        `json:"is_favorite"`
	FavoriteID  null.Uint64 `json:"-"`
}

type GetAllResponse struct {
	Items interface{} `json:"items"`
	Page  uint64      `json:"page"`
	Total uint64      `json:"total"`
}
