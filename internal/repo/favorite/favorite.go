package repository

import (
	"context"
	"database/sql"

	"github.com/lp/music-chart/connections/mysql"
	models "github.com/lp/music-chart/internal/model"
	repo "github.com/lp/music-chart/internal/repo"
)

const (
	RepoName = "Favorite Repository"

	Insert = `
		INSERT INTO
			favorite(user_id,music_id)
		VALUES
			(?,?)
	`

	Delete = `
		DELETE FROM
			favorite
		WHERE
			user_id = ?
			AND
			music_id = ?
	`
)

const (
	InsertNum = 10
	DeleteNum = 20
)

// Repository interface
type Repository interface {
	Insert(ctx context.Context, f models.Favorite) error
	Delete(ctx context.Context, userID, musicID uint64) (affectRow int64, err error)
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
			Query:     Insert,
			ActionNum: InsertNum,
		},
		repo.Stmt{
			Query:     Delete,
			ActionNum: DeleteNum,
		},
	}

	return &RepositoryImpl{
		db:   db,
		stmt: repo.BuildStmt(db, stmtList, RepoName),
	}
}

// Insert function
func (r *RepositoryImpl) Insert(ctx context.Context, f models.Favorite) error {
	args := []interface{}{
		f.UserID,
		f.MusicID,
	}

	_, err := r.stmt[InsertNum].ExecContext(ctx, args...)
	if err != nil {
		return err
	}

	return nil
}

// Delete function
func (r *RepositoryImpl) Delete(ctx context.Context, userID, musicID uint64) (affectRow int64, err error) {
	args := []interface{}{
		userID,
		musicID,
	}

	row, err := r.stmt[DeleteNum].ExecContext(ctx, args...)
	if err != nil {
		return 0, err
	}

	affectRow, _ = row.RowsAffected()
	return
}
