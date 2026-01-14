package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/nerdydave2017/filesharing/internal/infra/config"
)

func NewPostgresDB(config *config.Config) (*sql.DB, error) {
	// Format postgres db info
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword)

	// Open connection to db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Custom db settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Set connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping db and verify connection
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
