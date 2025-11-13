package main

import (
	"StudyApiServer/config"
	"StudyApiServer/internal/repository"
	"StudyApiServer/internal/router"
)

func main() {
	cfg := config.NewConfig()
	storage, err := repository.NewSQLiteStorage("database.db")
	if err != nil {
		panic(err)
	}

	r := router.NewRouter(
		storage,
		cfg.ApiToken,
	)

	r.Start()
}
