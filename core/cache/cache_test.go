package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestNewRedisCache(t *testing.T) {
	//redisConf := conf.GetRedisConf()
	//redisCache := NewRedisCache(redisConf)
	//ctx := context.Background()
	//redisCache.Store(ctx,"key","abc")
	cache := NewRedisCache("r-2ze2dfeae91a8e84.redis.rds.aliyuncs.com:6379", "Decrds2018", "prefix", 20, false)
	ctx := context.Background()
	err := cache.Store(ctx, "key", "value")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("111")
}
