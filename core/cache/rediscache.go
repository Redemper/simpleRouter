package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"simpleRouter/core/pojo"
	"time"
)

type RedisCache struct {
	client                *redis.Client
	Addr                  string
	Password              string
	Db                    int
	DefaultExpireDuration time.Duration
	EnableKeyPrefix       bool
	Prefix                string
}

func newRedisCache(conf *pojo.RedisConf) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
	})
	return &RedisCache{
		client:                rdb,
		Addr:                  conf.Addr,
		Password:              conf.Password,
		Db:                    conf.Db,
		DefaultExpireDuration: conf.DefaultExpireDuration,
		EnableKeyPrefix:       conf.EnableKeyPrefix,
		Prefix:                conf.Prefix,
	}
}

func (rc *RedisCache) Load(ctx context.Context, key string) (interface{}, error) {
	result, err := rc.client.Get(ctx, rc.Prefix+key).Result()
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (rc *RedisCache) Store(ctx context.Context, key string, value interface{}) error {
	_, err := rc.client.Set(ctx, rc.Prefix+key, value, rc.DefaultExpireDuration).Result()
	return err
}

func (rc *RedisCache) Delete(ctx context.Context, key string) error {
	_, err := rc.client.Del(ctx, rc.Prefix+key).Result()
	return err
}
