package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
    "time"
)

type RedisCache struct {
    client *redis.Client
    cf     *RedisConf
}

func NewRedisCache(conf *RedisConf) *RedisCache {
    rdb := redis.NewClient(&redis.Options{
        Addr:     conf.Addr,
        Password: conf.Password,
        DB:       conf.Db,
    })
    if !conf.EnableKeyPrefix {
        conf.Prefix = ""
    }
    return &RedisCache{
        client: rdb,
        cf:     conf,
    }
}

func (rc *RedisCache) Load(ctx context.Context, key string) (interface{}, error) {
    result, err := rc.client.Get(ctx, rc.cf.Prefix+key).Result()
    if err != nil {
        return nil, err
    } else {
        return result, nil
    }
}

func (rc *RedisCache) Store(ctx context.Context, key string, value interface{}) error {
    _, err := rc.client.Set(ctx, rc.cf.Prefix+key, value, rc.cf.DefaultExpireDuration).Result()
    return err
}

func (rc *RedisCache) Delete(ctx context.Context, key string) error {
    _, err := rc.client.Del(ctx, rc.cf.Prefix+key).Result()
    return err
}

type RedisConf struct {
    Addr                  string
    Password              string
    Db                    int
    DefaultExpireDuration time.Duration
    EnableKeyPrefix       bool
    Prefix                string
}
