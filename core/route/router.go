package route

import (
	"gorm.io/gorm"
	"sync"
)

type Router struct {
	Name      string `gorm:"name"`
	OriginUri string `gorm:"origin_uri"`
	Order     int    `gorm:"order"`
	Enabled   bool   `gorm:"enabled"`
	gorm.Model
}

var routerMap sync.Map

func AddRouter(r *Router) {
	routerMap.Store(r.Name, r)
}

func GetRouter(name string) *Router {
	route, ok := routerMap.Load(name)
	if !ok {
		return nil
	}
	realRouter, ok := route.(*Router)
	if ok {
		return realRouter
	}
	return nil
}

func GetAllRouter() []*Router {
	var routes []*Router
	routerMap.Range(func(key, value interface{}) bool {
		router, ok := value.(*Router)
		if ok {
			routes = append(routes, router)
		}
		return true
	})
	return routes
}
