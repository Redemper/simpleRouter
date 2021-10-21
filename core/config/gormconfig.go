package config

import (
    "gopkg.in/yaml.v2"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"
)

type DBConf struct {
    Dialector string `yaml:"dialector"`
    Url string `yaml:"url"`
    User string `yaml:"user"`
    Password string `yaml:"password"`
    Port string `yaml:"port"`
}

var iDB gorm.DB

func init(){
    // read conf
    conf := new(DBConf)
    path, errp := filepath.Abs("../../conf")
    if errp == nil {
        yamlFile, err := ioutil.ReadFile(path + "/gorm.yml")
        if err != nil {
            log.Fatalf("yamlFile.Get err #%v ", err)
        }
        err = yaml.Unmarshal(yamlFile, conf)
        // err = yaml.Unmarshal(yamlFile, &resultMap)
        if err != nil {
            log.Fatalf("Unmarshal: %v", err)
        }
        //db, err := gorm.Open(mysql.Open(conf.generateDBString()), &gorm.Config{})

        //iDB = gorm.Open(mysql.open(generateDBString()))
    }
}

func (conf *DBConf) generateDBString() string{
    //var conStr = "cd_mall:Cd_Mall1@tcp(rm-2zek85dzv7g624hbz.mysql.rds.aliyuncs.com:3306)/mall_shop?charset=utf8mb4&parseTime=True&loc=Local"
    var result strings.Builder
    result.WriteString(conf.User)
    result.WriteString(":")
    result.WriteString(conf.Password)
    result.WriteString("@tcp(")
    result.WriteString(conf.Url)
    result.WriteString(":")
    result.WriteString(conf.Port)
    result.WriteString(")/mall_shop?charset=utf8mb4&parseTime=True&loc=Local")
    return result.String()
}