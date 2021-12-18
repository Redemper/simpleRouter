package cache

import (
	"context"
	"simpleRouter/core/pojo"
)

type ICache interface {
	Load(ctx context.Context, key string) (interface{}, error)
	Store(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

var c ICache

func InitCacheByCacheType(cacheType string, conf *pojo.RedisConf) error {
	switch cacheType {
	case "redis":
		c = newRedisCache(conf)
	case "map":
		c = newMemoryCache(MAP)
	case "syncMap":
		fallthrough
	default:
		c = newMemoryCache(SYNCMAP)
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
