package cache

import (
	"context"
	"fmt"
	"simpleRouter/core/pojo"
	"testing"
)

func TestNewRedisCache(t *testing.T) {
	//redisConf := conf.GetRedisConf()
	//redisCache := newRedisCache(redisConf)
	//ctx := context.Background()
	//redisCache.Store(ctx,"key","abc")
	conf := new(pojo.RedisConf)
	conf.Addr = "r-2ze2dfeae91a8e84.redis.rds.aliyuncs.com:6379"
	conf.Db = 20
	conf.Password = "Decrds2018"
	conf.Prefix = "pre"
	conf.DefaultExpireDuration = 10
	conf.EnableKeyPrefix = false
	cache := newRedisCache(conf)
	ctx := context.Background()
	err := cache.Store(ctx, "key", "value")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("111")
}
