package pojo

type ServerConf struct {
	GraceShutDown bool      `yaml:"grace-shutdown"`
	ServerPort    string    `yaml:"server-port"`
	Rl            RateLimit `yaml:"rateLimit"`
}

type RateLimit struct {
	Enable bool `yaml:"enable"`
	Rps    int  `yaml:"rps"`
}
