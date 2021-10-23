/**
  load router config from yaml
*/
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

func InitRouterFromYaml() {
	path, errp := filepath.Abs("../../conf")
	log.Println("start read router config from yaml")
	yr := new(yamlRouters)
	if errp == nil {
		yamlFile, err := ioutil.ReadFile(path + "/routers.yml")
		if err != nil {
			log.Printf("initRouterFromYaml yamlFile.Get err #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, yr)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		routers := yr.Routers
		for _, r := range routers {
			AddRouter(r)
		}
	}
}

type yamlRouters struct {
	Routers []*Router `yaml:"router"`
}
