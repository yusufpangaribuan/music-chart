package repository

import (
	"context"
	"database/sql"

	"github.com/lp/music-chart/connections/mysql"
	models "github.com/lp/music-chart/internal/model"
	repo "github.com/lp/music-chart/internal/repo"
)

const (
	RepoName      = "User Repository"
	GetByUserName = `
		SELECT 
			id, username, password 
		FROM 
			user
		WHERE username=?
	`

	AddUser = `
		INSERT INTO
			user(full_name,username,password,gender,hobby,address)
		VALUES
			(?,?,?,?,?,?)
	`
)

const (
	GetByUserNameNum = 10
	AddUserNum       = 20
)

// Repository interface
type Repository interface {
	GetByUserName(ctx context.Context, userName string) (*models.User, error)
	Insert(ctx context.Context, user models.User) (id uint64, err error)
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
			Query:     GetByUserName,
			ActionNum: GetByUserNameNum,
		},
		repo.Stmt{
			Query:     AddUser,
			ActionNum: AddUserNum,
		},
	}

	return &RepositoryImpl{
		db:   db,
		stmt: repo.BuildStmt(db, stmtList, RepoName),
	}
}

// GetByUserName function
func (r *RepositoryImpl) GetByUserName(ctx context.Context, userName string) (*models.User, error) {
	row := r.stmt[GetByUserNameNum].QueryRowContext(ctx, userName)

	var user models.User
	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert function
func (r *RepositoryImpl) Insert(ctx context.Context, user models.User) (id uint64, err error) {
	args := []interface{}{
		user.FullName,
		user.UserName,
		user.Password,
		user.Gender,
		user.Hobby,
		user.Address,
	}

	row, err := r.stmt[AddUserNum].ExecContext(ctx, args...)
	if err != nil {
		return
	}

	id64, _ := row.LastInsertId()
	id = uint64(id64)
	return
}
