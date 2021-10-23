package conf

import "path/filepath"

func InitServerConf() (*ServerConf, error) {
	sc := new(ServerConf)
	path, errp := filepath.Abs("../../conf")
	if errp != nil {
		return nil, errp
	}
	err := ReadYaml(path+"/server.yml", sc)
	if err != nil {
		return nil, err
	}
	return sc, nil
}

type ServerConf struct {
	GraceShutDown bool      `yaml:"grace-shutdown"`
	ServerPort    string    `yaml:"server-port"`
	Rl            RateLimit `yaml:"rateLimit"`
}

type RateLimit struct {
	Enable bool `yaml:"enable"`
	Rps    int  `yaml:"rps"`
}
