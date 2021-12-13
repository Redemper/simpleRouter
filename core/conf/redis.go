package conf

import "time"

type RedisConf struct {
	Addr                  string        `yaml:"addr"`
	Password              string        `yaml:"password"`
	Db                    int           `yaml:"db"`
	DefaultExpireDuration time.Duration `yaml:"expire"`
	EnableKeyPrefix       bool          `yaml:"enable_key_prefix"`
	Prefix                string        `yaml:"prefix"`
}
