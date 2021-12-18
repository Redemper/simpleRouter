package cache

import (
	"log"
	"simpleRouter/core/conf"
	"simpleRouter/core/pojo"
)

// 初始化缓存
func init() {
	initCache()
}

func initCache() {
	// 读取yaml配置到对应的实体中
	redisConf := new(pojo.RedisConf)
	err := conf.ReadYamlFromDefaultPath(redisConf)
	if err != nil {
		log.Fatal("can load redis conf from yaml")
		InitCacheByCacheType("map", nil)
	} else {
		InitCacheByCacheType("redis", redisConf)
	}
}
