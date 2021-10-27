/**
  load router config from yaml
*/
package route

import (
	"log"
	"path/filepath"
	"simpleRouter/core/conf"
)

func InitRouterFromYaml() []*Router {
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
		return routers
	}
	return nil
}

type yamlRouters struct {
	Routers []*Router `yaml:"router"`
}