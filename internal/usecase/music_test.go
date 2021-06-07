package usecase

import (
	"context"
	"errors"
	"testing"

	models "github.com/lp/music-chart/internal/model"
	r "github.com/lp/music-chart/internal/repo/music"
	favorite "github.com/lp/music-chart/internal/repo/music/mocks"
	music "github.com/lp/music-chart/internal/repo/music/mocks"
	user "github.com/lp/music-chart/internal/repo/music/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
)

var mockMusic *music.Repository
var mockFavorite *favorite.Repository
var mockUser *user.Repository

func setup() {
	mockMusic = new(music.Repository)
	mockFavorite = new(favorite.Repository)
	mockUser = new(user.Repository)
}

func Test_GetAll(t *testing.T) {
	mockErr := errors.New("error mock")
	setup()
	type fields struct {
		MusicRepo r.Repository
	}
	type args struct {
		ctx    context.Context
		params models.BasicSelectParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		patch   func(args args)
		want    *models.GetAllResponse
		wantErr bool
	}{
		{
			name: "When_GetuserInfoFromContextReturnError_Expect_Error",
			fields: fields{
				MusicRepo: mockMusic,
			},
			args: args{
				ctx:    context.TODO(),
				params: models.BasicSelectParams{},
			},
			patch: func(args args) {
				handleGetUserInfo = func(ctx context.Context) (*models.UserInfo, error) {
					return nil, mockErr
				}
			},
			wantErr: true,
		},
		{
			name: "When_GetAllReturnError_Expect_Error",
			fields: fields{
				MusicRepo: mockMusic,
			},
			args: args{
				ctx:    context.TODO(),
				params: models.BasicSelectParams{},
			},
			patch: func(args args) {
				handleGetUserInfo = func(ctx context.Context) (*models.UserInfo, error) {
					return &models.UserInfo{
						ID: 1,
					}, nil
				}
				mockMusic.On("GetAll", args.ctx, args.params.Limit, args.params.Page).Return(nil, uint64(0), mockErr).Once()
			},
			wantErr: true,
		},
		{
			name: "When_GetAllReturnSuccess_Expect_Success",
			fields: fields{
				MusicRepo: mockMusic,
			},
			args: args{
				ctx:    context.TODO(),
				params: models.BasicSelectParams{},
			},
			patch: func(args args) {
				handleGetUserInfo = func(ctx context.Context) (*models.UserInfo, error) {
					return &models.UserInfo{
						ID: 1,
					}, nil
				}
				mockMusic.On("GetAll", args.ctx, args.params.Limit, args.params.Page).Return([]*models.Music{
					&models.Music{
						ID:         1,
						FavoriteID: null.NewUint64(1, true),
					},
				}, uint64(0), nil).Once()
			},
			want: &models.GetAllResponse{
				Items: []*models.Music{
					&models.Music{
						ID:         1,
						IsFavorite: true,
						FavoriteID: null.NewUint64(1, true),
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := MusicUsecaseImpl{
				MusicRepo: tt.fields.MusicRepo,
			}
			tt.patch(tt.args)
			got, err := c.GetAll(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}

}
