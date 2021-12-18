package pojo

type YamlConf struct {
	*DBConf     `yaml:"db"`
	*NacosConf  `yaml:"nacos"`
	*ServerConf `yaml:"server"`
	Routers     []*Router `yaml:"yaml_router"`
	CacheType   string    `yaml:"cache_type"`
	*LbConf     `yaml:"lb"`
	*RedisConf  `yaml:"redis"`
}
