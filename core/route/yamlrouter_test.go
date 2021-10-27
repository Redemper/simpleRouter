package route

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestInitRouterFromYaml(t *testing.T) {
	path, errp := filepath.Abs("../../../conf")
	yr := new(yamlRouters)
	if errp == nil {
		fmt.Println("path === ", path)
		yamlFile, err := ioutil.ReadFile(path + "/routers.yml")

		// conf := new(module.Yaml1)
		// yamlFile, err := ioutil.ReadFile("test.yaml")

		// conf := new(module.Yaml2)
		//  yamlFile, err := ioutil.ReadFile("test1.yaml")

		log.Println("yamlFile:", string(yamlFile))
		if err != nil {
			log.Printf("yamlFile.Get err #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, yr)
		// err = yaml.Unmarshal(yamlFile, &resultMap)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		log.Println("yr", *yr)
		log.Println("yr", yr.Routers[0])
		// log.Println("conf", resultMap)
	}
}
