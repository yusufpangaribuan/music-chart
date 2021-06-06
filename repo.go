package main

import (
	"log"
	"reflect"
	"sync"
	"time"

	conMysql "github.com/lp/music-chart/connections/mysql"

	favoriteRepo "github.com/lp/music-chart/internal/repo/favorite"
	musicRepo "github.com/lp/music-chart/internal/repo/music"
	userRepo "github.com/lp/music-chart/internal/repo/user"
)

type repo struct {
	music    musicRepo.Repository
	user     userRepo.Repository
	favorite favoriteRepo.Repository
}

var globalConn conMysql.DB

func initRepo() {
	globalConn = conMysql.DB{
		Client: conMysql.GetSQLClient(),
	}

	//variable repos already on app.go
	var wg sync.WaitGroup
	totalReposField := reflect.ValueOf(repos).NumField()
	wg.Add(totalReposField)

	go func() {
		defer wg.Done()
		start := time.Now()
		repos.music = initRepoMusic()
		log.Printf("Finish initRepoMusic() in %v", time.Since(start).Seconds())
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		repos.user = initRepoUser()
		log.Printf("Finish initRepoUser() in %v", time.Since(start).Seconds())
	}()

	go func() {
		defer wg.Done()
		start := time.Now()
		repos.favorite = initRepoFavorite()
		log.Printf("Finish initRepoFavorite() in %v", time.Since(start).Seconds())
	}()

	wg.Wait()

}

func initRepoMusic() musicRepo.Repository {
	return musicRepo.NewRepositoryImpl(globalConn)
}

func initRepoUser() userRepo.Repository {
	return userRepo.NewRepositoryImpl(globalConn)
}

func initRepoFavorite() favoriteRepo.Repository {
	return favoriteRepo.NewRepositoryImpl(globalConn)
}
