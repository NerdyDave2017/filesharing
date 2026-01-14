package main

import (
	"database/sql"

	"github.com/nerdydave2017/filesharing/internal/infra/cache"
	"github.com/nerdydave2017/filesharing/internal/infra/config"
	"github.com/nerdydave2017/filesharing/internal/infra/db"
	"github.com/redis/go-redis/v9"
)

type Application struct {
	Config *config.Config
	DB     *sql.DB
	Cache  *redis.Client
}

func App() (*Application, error) {
	app := &Application{}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	app.Config = cfg

	// Create DB instance
	db, err := db.NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}
	app.DB = db

	// Create Redis client
	cache, err := cache.NewRedisClient(cfg)
	if err != nil {
		return nil, err
	}
	app.Cache = cache

	return app, nil

}

func (app *Application) Close() error {
	if app.DB != nil {
		app.DB.Close()
	}

	if app.Cache != nil {
		app.Cache.Close()
	}

	return nil
}
