/**
  load router config from yaml
*/
package config

import (
	"log"
	"path/filepath"
	"simpleRouter/core/conf"
)

func InitRouterFromYaml() {
	path, errp := filepath.Abs("../../conf")
	log.Println("start read router config from yaml")
	yr := new(yamlRouters)
	if errp != nil {
		panic(errp)
	}
	err := conf.ReadYaml(path+"/routers.yml", yr)
	if err != nil {
		panic(err)
	} else {
		routers := yr.Routers
		for _, r := range routers {
			AddRouter(r)
		}
	}
}

type yamlRouters struct {
	Routers []*Router `yaml:"router"`
}
