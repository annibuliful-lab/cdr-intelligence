package clients

import (
	"backend/src/config"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

// NewRedisClient ensures thread-safe initialization of the Redis client.
func NewRedisClient() (*redis.Client, error) {
	var err error
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.GetEnv("REDIS_HOST", "localhost:6379"),
			Password: config.GetEnv("REDIS_PASSWORD", ""),
			DB:       0,
		})
	})

	if err != nil {
		return nil, err
	}

	return redisClient, nil
}
