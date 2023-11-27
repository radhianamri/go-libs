package redis

import (
	"context"
	"log"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type Cache struct {
	Cache *redis.Client
}

type Config struct {
	Host           string
	Password       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxRetries     int
	MaxActiveConns int
}

func Init(config Config) *Cache {
	return &Cache{
		Cache: redis.NewClient(&redis.Options{
			Addr:         config.Host,
			Password:     config.Password,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
			MaxRetries:   config.MaxRetries,
		}),
	}
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := c.Cache.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		log.Printf("Failed to get redis cache key %s: %s\n", key, err.Error())
		return nil, err
	}
	return []byte(res), nil
}

func (c *Cache) Set(ctx context.Context, key string, data []byte, ttl time.Duration) error {
	_, err := c.Cache.Set(ctx, key, string(data), ttl).Result()
	if err != nil {
		log.Printf("Failed to set redis cache key %s: %s\n", key, err.Error())
		return err
	}
	return nil
}
