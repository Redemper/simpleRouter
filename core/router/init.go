package router

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simpleRouter/core/conf"
	"simpleRouter/core/pojo"
)

// 从db或者yaml中读取router

func init() {
	initRouter()
}

func initRouterFromYaml() []*pojo.Router {
	routers := GetYamlRouters()
	if nil == routers {
		return nil
	}
	return routers
}

func GetYamlRouters() []*pojo.Router {
	//var routers YamlRouters
	routers := conf.YamlConf.Routers
	return routers
}

type YamlRouters struct {
	Routers []*pojo.Router `yaml:"router"`
}

func initRouterFromDB() []*pojo.Router {
	db, err := getDb()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&pojo.Router{})
	var routes []*pojo.Router
	db.Where(" enabled = 1 ").Find(&routes)
	return routes
}

func getDb() (*gorm.DB, error) {
	dbConf := conf.YamlConf.DBConf
	dbString := dbConf.GenerateDBString()
	return gorm.Open(mysql.Open(dbString), &gorm.Config{})
}
