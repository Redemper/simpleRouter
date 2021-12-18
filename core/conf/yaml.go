package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

const DEFAULT_YAML_PATH = "../../conf"

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

func ReadYamlFromDefaultPath(out interface{}) error {
	path, pathErr := filepath.Abs(DEFAULT_YAML_PATH)
	if pathErr != nil {
		return errors.New("can find path")
	}
	return ReadYaml(path, out)
}
