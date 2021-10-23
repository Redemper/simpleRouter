package conf

import (
    "gopkg.in/yaml.v2"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"
)

/**
  db conf.
  Url : 地址
  Port : 端口
  User : 用户名
  Password : 密码
  Schema : 数据库schema
*/
type DBConf struct {
    Dialector string `yaml:"dialector"`
    Url       string `yaml:"url"`
    User      string `yaml:"user"`
    Password  string `yaml:"password"`
    Port      string `yaml:"port"`
    Schema    string `yaml:"schema"`
}

func (conf *DBConf) GenerateDBString() string {
    //var conStr = "cd_mall:Cd_Mall1@tcp(rm-2zek85dzv7g624hbz.mysql.rds.aliyuncs.com:3306)/mall_shop?charset=utf8mb4&parseTime=True&loc=Local"
    var result strings.Builder
    result.WriteString(conf.User)
    result.WriteString(":")
    result.WriteString(conf.Password)
    result.WriteString("@tcp(")
    result.WriteString(conf.Url)
    result.WriteString(":")
    result.WriteString(conf.Port)
    result.WriteString(")/")
    result.WriteString(conf.Schema)
    result.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
    return result.String()
}

func GetDBFromYaml() (*gorm.DB,error){
    // read conf
    conf := new(DBConf)
    path, errp := filepath.Abs("../../conf")
    if errp != nil {
        return nil, errp
    }
    yamlFile, err := ioutil.ReadFile(path + "/gorm.yml")
    if err != nil {
        log.Fatalf("yamlFile.Get err #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, conf)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    dbString := conf.GenerateDBString()
    db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})

    return db,err
}
