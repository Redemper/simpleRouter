package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadYaml(path string, out interface{}) error {
	log.Printf("start read yaml,path == %v", path)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("yamlFile.Get err #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, out)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}
