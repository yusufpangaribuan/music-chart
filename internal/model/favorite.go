package models

// Favorite struct
type Favorite struct {
	UserID  uint64
	MusicID uint64
}

// SetFavoriteResponse struct
type SetFavoriteResponse struct {
	IsFavorite bool `json:"is_favorite"`
}
