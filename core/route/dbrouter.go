/**
load router from db
*/
package route

import (
	"gorm.io/gorm"
	"simpleRouter/core/conf"
)

var iDB gorm.DB

func InitRouterFromDB() []*Router {
	db, err := conf.GetDBFromYaml()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Router{})
	var routes []*Router
	db.Where(" enabled = 1 ").Find(&routes)
	return routes
}
