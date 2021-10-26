/**
load router from db
*/
package config

import (
	"gorm.io/gorm"
	"simpleRouter/core/conf"
)

var iDB gorm.DB

func GetDB() gorm.DB {
	return iDB
}

func InitRouterFromDB() []*Router {
	db, err := conf.GetDBFromYaml()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Router{})
	var routes []*Router
	db.Where(" enabled = 1 ").Find(&routes)
	//for _, r := range routes {
	//	// init route
	//	route.AddRouter(r)
	//}
	return routes
}
