package conf

type NacosClient struct {
	NameSpaceId          string `yaml:"namespace"`
	TimeoutMs            uint64 `yaml:"timeout-ms"`
	BeatInterval         uint64 `yaml:"beat-interval"`
	OpenKMS              bool   `yaml:"open-kms"`
	CacheDir             string `yaml:"cache-dir"`
	UpdateThreadNum      uint64 `yaml:"update-thread-num"`
	NotLoadCacheAtStart  bool   `yaml:"not-load-cache-start"`
	UpdateCacheWhenEmpty bool   `yaml:"update-cache-when-empty"`
	LogDir               string `yaml:"log-dir"`
	RotateTime           string `yaml:"rotate-time"`
	MaxAge               uint64 `yaml:"max-age"`
	LogLevel             string `yaml:"log-level"`
	UserName             string `yaml:"username"`
	Password             string `yaml:"password"`
}

type NacosServer struct {
	Scheme      string `yaml:"scheme"`
	ContextPath string `yaml:"context-path"`
	IpAddr      string `yaml:"ip-addr"`
	Port        uint64 `yaml:"port"`
}

type NacosConf struct {
	Cc NacosClient   `yaml:"client-conf"`
	Sc []NacosServer `yaml:"server-conf"`
}

func ReadNacosConfFromYaml() *NacosConf {
	// read conf
	nacosConf := GetNacosConf()
	return nacosConf
}
