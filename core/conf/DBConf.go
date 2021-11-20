package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

// generate gorm db link.
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

/**
get db config from yaml
*/
func GetDBFromYaml() (*gorm.DB, error) {
	// read conf
	conf := GetDbConf()
	dbString := conf.GenerateDBString()
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})

	return db, err
}
