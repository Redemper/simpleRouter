package cache

import (
	"context"
	"flag"
)

type ICache interface {
	Load(ctx context.Context, key string) (interface{}, error)
	Store(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

var CacheType = flag.String("choose your caceh type", "memory", "")

var c ICache

func init() {
	flag.Parse()
	switch *CacheType {
	case "redis":
	case "memory":
	default:
		c = NewMemoryCache(SYCNMAP)
	}
}

func Load(ctx context.Context, key string) (interface{}, error) {
	return c.Load(ctx, key)
}

func LoadWithoutContext(key string) (interface{}, error) {
	return c.Load(context.Background(), key)
}

func Store(ctx context.Context, key string, value interface{}) error {
	return c.Store(ctx, key, value)
}

func StoreWithoutContext(key string, value interface{}) error {
	return c.Store(context.Background(), key, value)
}

func Delete(ctx context.Context, key string) error {
	return c.Delete(ctx, key)
}

func DeleteWithoutContext(key string) error {
	return c.Delete(context.Background(), key)
}
