package main

import (
	"log"
	"reflect"
	"sync"
	"time"

	u "github.com/lp/music-chart/internal/usecase"
)

type usecase struct {
	music    u.MusicUsecase
	favorite u.FavoriteUsecase
	user     u.Auth
}

func initUsecase() {
	//variable repos already on app.go
	var wg sync.WaitGroup
	totalUsecasesField := reflect.ValueOf(usecases).NumField()
	wg.Add(totalUsecasesField)

	go func() {
		defer wg.Done()
		start := time.Now()
		usecases.music = initUsecaseMusic()
		log.Printf("Finish initUsecaseMusic() in %v", time.Since(start).Seconds())
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		usecases.user = initUsecaseUser()
		log.Printf("Finish initUsecaseUser() in %v", time.Since(start).Seconds())
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		usecases.favorite = initUsecaseFavorite()
		log.Printf("Finish initAuth() in %v", time.Since(start).Seconds())
	}()

	wg.Wait()

}

func initUsecaseMusic() u.MusicUsecase {
	return u.NewMusicUsecaseImpl(u.MusicUsecaseImpl{
		MusicRepo: repos.music,
	})
}

func initUsecaseUser() u.Auth {
	return u.NewAuthImpl(u.AuthImpl{
		UserRepo: repos.user,
	})
}

func initUsecaseFavorite() u.FavoriteUsecase {
	return u.NewFavoriteUsecaseImpl(u.FavoriteUsecaseImpl{
		FavoriteRepo: repos.favorite,
		MusicRepo:    repos.music,
	})
}
