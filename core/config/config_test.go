package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestReadYaml(t *testing.T) {
	// resultMap := make(map[string]interface{})
	conf := new(DBConf)
	path, errp := filepath.Abs("../../conf")
	if errp == nil {
		fmt.Println("path === ", path)
		yamlFile, err := ioutil.ReadFile(path + "/gorm.yml")

		// conf := new(module.Yaml1)
		// yamlFile, err := ioutil.ReadFile("test.yaml")

		// conf := new(module.Yaml2)
		//  yamlFile, err := ioutil.ReadFile("test1.yaml")

		log.Println("yamlFile:", yamlFile)
		if err != nil {
			log.Printf("yamlFile.Get err #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		// err = yaml.Unmarshal(yamlFile, &resultMap)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		log.Println("conf", conf)
		// log.Println("conf", resultMap)
	}
}

func TestGetAllRouter(t *testing.T) {
	routers := GetAllRouter()
	for _, r := range routers {
		marshal, err := json.Marshal(r)
		if err == nil {
			fmt.Println(string(marshal))
		} else {
			fmt.Println("sry json fail")
		}
	}
}
