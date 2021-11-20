package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
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

var Yc *YamlConf = new(YamlConf)

type YamlConf struct {
	*DBConf     `yaml:"db"`
	*NacosConf  `yaml:"nacos"`
	*ServerConf `yaml:"server"`
}

func init() {
	path, pathErr := filepath.Abs("../../conf")
	if pathErr != nil {
		return
	}
	err := ReadYaml(path+"/prop.yml", Yc)
	if err != nil {
		return
	}
}

func GetDbConf() *DBConf {
	return Yc.DBConf
}

func GetNacosConf() *NacosConf {
	return Yc.NacosConf
}

func GetServerConf() *ServerConf {
	return Yc.ServerConf
}
