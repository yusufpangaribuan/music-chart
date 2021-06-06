package repository

import (
	"context"
	"database/sql"

	"github.com/lp/music-chart/connections/mysql"
	models "github.com/lp/music-chart/internal/model"
	repo "github.com/lp/music-chart/internal/repo"
)

const (
	RepoName = "Music Repository"
	GetByID  = `
		SELECT 
			m.id, m.title, m.singer, m.duration, m.album, m.release_year, f.user_id
		FROM 
			music m
		LEFT JOIN 
			favorite f
		ON 
			m.id = f.music_id
		WHERE m.id=?
	`

	GetAll = `
		SELECT 
			m.id, m.title, f.user_id
		FROM 
			music m
		LEFT JOIN 
			favorite f
		ON 
			m.id = f.music_id
		LIMIT ?
		OFFSET ?
	`

	GetTotal = `
		SELECT 
			COUNT(id)
		FROM 
			music
	`
)

const (
	GetByIDNum  = 10
	GetAllNum   = 20
	GetTotalNum = 30
)

// Repository interface
type Repository interface {
	GetByID(ctx context.Context, id uint64) (*models.Music, error)
	GetAll(ctx context.Context, limit, offset uint64) (res []*models.Music, tot uint64, err error)
}

// RepositoryImpl struct
type RepositoryImpl struct {
	db   mysql.DB
	stmt map[int]*sql.Stmt
}

// NewRepositoryImpl function
func NewRepositoryImpl(db mysql.DB) Repository {
	stmtList := []repo.Stmt{
		repo.Stmt{
			Query:     GetByID,
			ActionNum: GetByIDNum,
		},
		repo.Stmt{
			Query:     GetAll,
			ActionNum: GetAllNum,
		},
		repo.Stmt{
			Query:     GetTotal,
			ActionNum: GetTotalNum,
		},
	}

	return &RepositoryImpl{
		db:   db,
		stmt: repo.BuildStmt(db, stmtList, RepoName),
	}
}

// GetByID function
func (r *RepositoryImpl) GetByID(ctx context.Context, id uint64) (*models.Music, error) {
	row := r.stmt[GetByIDNum].QueryRowContext(ctx, id)

	var m models.Music
	err := row.Scan(
		&m.ID,
		&m.Title,
		&m.Singer,
		&m.Duration,
		&m.Album,
		&m.ReleaseYear,
		&m.FavoriteID,
	)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// GetAll function
func (r *RepositoryImpl) GetAll(ctx context.Context, limit, offset uint64) (res []*models.Music, tot uint64, err error) {
	rows, err := r.stmt[GetAllNum].QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var m models.Music
		err = rows.Scan(
			&m.ID,
			&m.Title,
			&m.FavoriteID,
		)
		if err != nil {
			return nil, 0, err
		}

		res = append(res, &m)
	}

	row := r.stmt[GetTotalNum].QueryRowContext(ctx)
	err = row.Scan(
		&tot,
	)
	if err != nil {
		return nil, 0, err
	}

	return
}
