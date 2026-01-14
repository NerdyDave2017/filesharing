package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/nerdydave2017/filesharing/internal/infra/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *config.Config) (*redis.Client, error) {
	// Format redis connection address
	addr := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)

	// Create redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     20,
		MinIdleConns: 2,
	})

	// Set connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping redis db and verify connection
	if err := rdb.Ping(ctx).Err(); err != nil {

		return nil, err
	}

	return rdb, nil
}
