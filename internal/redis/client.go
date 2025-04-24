package redisclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/noornee/toolkits/internal/environment"
	"github.com/noornee/toolkits/internal/model/env"
	"github.com/redis/go-redis/v9"
)

// Store defines methods for interacting with a Redis key-value store.
type Store interface {
	SetValue(ctx context.Context, key string, value any, ttl time.Duration) error
	GetValue(ctx context.Context, key string) (string, error)
}

// Redis is a wrapper around the Redis client with configuration.
type Redis struct {
	env    *environment.Env
	client *redis.Client
}

// NewRedisClient creates and returns a new Redis client after testing the connection.
func NewRedisClient(e *environment.Env) (Store, error) {
	var tlsConfig *tls.Config
	if !strings.EqualFold(e.Get(env.RedisTLSEnabled), "false") {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	redisDB := 0
	redisDBStr := e.Get(env.RedisDB)
	if !strings.EqualFold(redisDBStr, "") {
		if num, err := strconv.ParseInt(redisDBStr, 10, 64); err == nil {
			redisDB = int(num)
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:      e.Get(env.RedisAddr),
		Password:  e.Get(env.RedisPassword),
		Username:  e.Get(env.RedisUsername),
		DB:        redisDB,
		TLSConfig: tlsConfig,
	})

	// Use a context with timeout for pinging Redis
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("Redis connection failed: %v", err)
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Redis{
		env:    e,
		client: client,
	}, nil
}

// SetValue stores a key-value pair in Redis with a TTL.
func (r *Redis) SetValue(ctx context.Context, key string, value any, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}

// GetValue retrieves the value of a key from Redis.
func (r *Redis) GetValue(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
