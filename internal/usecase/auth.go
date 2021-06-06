package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	helper "github.com/lp/music-chart/internal/helper"
	models "github.com/lp/music-chart/internal/model"
	userRepo "github.com/lp/music-chart/internal/repo/user"
	"golang.org/x/crypto/bcrypt"
)

// Auth interface
type Auth interface {
	Login(ctx context.Context, req models.LoginReq) (token string, httpCode int, err error)
	Register(ctx context.Context, req models.User) (token string, httpCode int, err error)
}

// AuthImpl struct
type AuthImpl struct {
	UserRepo userRepo.Repository
}

// NewAuthImpl function
func NewAuthImpl(auth AuthImpl) Auth {
	return &auth
}

// Login function
func (a *AuthImpl) Login(ctx context.Context, req models.LoginReq) (token string, httpCode int, err error) {
	httpCode = http.StatusOK
	user, err := a.UserRepo.GetByUserName(ctx, req.UserName)
	if err == sql.ErrNoRows {
		err = errors.New("account is not exist")
		return
	}

	if err != nil {
		return
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if hashErr != nil {
		return "", http.StatusOK, errors.New("wrong email or password")
	}

	token, err = helper.GetToken(user)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return
}

// Register function
func (a *AuthImpl) Register(ctx context.Context, req models.User) (token string, httpCode int, err error) {
	httpCode = http.StatusOK
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	user := models.User{
		FullName: req.FullName,
		UserName: req.UserName,
		Password: string(hashedPassword),
		Gender:   req.Gender,
		Hobby:    req.Hobby,
		Address:  req.Address,
	}

	id, err := a.UserRepo.Insert(ctx, user)
	me, ok := err.(*mysql.MySQLError)
	if !ok {
		user.ID = id
		token, err = helper.GetToken(&user)
		if err != nil {
			return "", http.StatusInternalServerError, err
		}
		return
	}

	if me.Number == 1062 {
		return "", http.StatusNotAcceptable, errors.New("username is already exists in a database.")
	}

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return
}
