package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
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

func NewRedisCache(Addr, Password, Prefix string, Db int, EnableKeyPrefix bool) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       Db,
	})
	if !EnableKeyPrefix {
		Prefix = ""
	}
	return &RedisCache{
		client: rdb,
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
