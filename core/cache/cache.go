package cache

import (
	"context"
)

type ICache interface {
	Load(ctx context.Context, key string) (interface{}, error)
	Store(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

var c ICache

func InitCacheByCacheType(cacheType string, m map[string]interface{}) error {
	switch cacheType {
	case "redis":
		addr := m["Addr"].(string)
		password := m["Password"].(string)
		Db := m["Db"].(int)
		Prefix := m["prefix"].(string)
		enable := m["EnableKeyPrefix"].(bool)
		c = NewRedisCache(addr, password, Prefix, Db, enable)
	default:
		c = NewMemoryCache(SYCNMAP)
	}
	//TODO 添加一些错误
	return nil
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
