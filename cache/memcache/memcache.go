package memcache

import (
	"context"
	"errors"
	"time"

	"github.com/karlseguin/ccache/v3"
)

type Cache struct {
	Cache   *ccache.Cache[[]byte]
	GetFunc func(key string) *ccache.Item[[]byte]
}

type Config struct {
	MaxSize     int64
	WithPromote bool
}

func Init(config Config) *Cache {
	c := ccache.Configure[[]byte]()
	if config.MaxSize > 0 {
		c.MaxSize(config.MaxSize)
	}

	cache := &Cache{
		Cache: ccache.New[[]byte](c),
	}
	cache.initGetFunc(config.WithPromote)

	return cache
}

func (c *Cache) initGetFunc(withPromote bool) {
	c.GetFunc = func(key string) *ccache.Item[[]byte] {
		if withPromote {
			return c.Cache.Get(key)
		}
		return c.Cache.GetWithoutPromote(key)
	}
}

var errorNil = errors.New("memcache: nil")

func (c *Cache) Get(_ context.Context, key string) ([]byte, error) {
	item := c.GetFunc(key)
	if item == nil || item.Expired() {
		return nil, errorNil
	}

	return item.Value(), nil
}

func (c *Cache) Set(_ context.Context, key string, data []byte, ttl time.Duration) error {
	c.Cache.Set(key, data, ttl)
	return nil
}
