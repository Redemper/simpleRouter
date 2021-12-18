package cache

import "simpleRouter/core/conf"

func init() {
	redisConf := conf.GetRedisConf()
	if nil == redisConf {
		return
	}
	m := make(map[string]interface{})
	m["Addr"] = redisConf.Addr
	m["Password"] = redisConf.Password
	m["Db"] = redisConf.Db
	m["prefix"] = redisConf.Prefix
	m["EnableKeyPrefix"] = redisConf.EnableKeyPrefix
	InitCacheByCacheType("redis", m)
}
