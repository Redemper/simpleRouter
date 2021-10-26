package route

import (
	"flag"
	"log"
	"simpleRouter/core/route/config"
	"sync"
)

var routeLoadType = flag.String("route-load-type", "db", "the way to load routers,Optional value:db,yaml")

var RouterMap sync.Map

func init() {
	flag.Parse()
	log.Println("routeLoadType is ", routeLoadType)
	var routers []*config.Router
	switch *routeLoadType {
	case "yaml":
		routers = config.InitRouterFromYaml()
	case "db":
		fallthrough
	default:
		routers = config.InitRouterFromDB()
	}
	log.Println("routeLoad finished,routes :")
	for _, r := range routers {
		AddRouter(r)
	}
}

var uriRouterMap sync.Map
var nameRouterMap sync.Map

func AddRouter(r *config.Router) {
	uriRouterMap.Store(r.OriginUri, r)
	nameRouterMap.Store(r.Name, r)
}

func GetAllRouter() []*config.Router {
	var routes []*config.Router
	uriRouterMap.Range(func(key, value interface{}) bool {
		router, ok := value.(*config.Router)
		if ok {
			routes = append(routes, router)
		}
		return true
	})
	return routes
}

func GetRouterByUri(uri string) *config.Router {
	route, ok := uriRouterMap.Load(uri)
	if !ok {
		return nil
	}
	realRouter, ok := route.(*config.Router)
	if ok {
		return realRouter
	}
	return nil
}

func GetRouterByName(name string) *config.Router {
	route, ok := nameRouterMap.Load(name)
	if !ok {
		return nil
	}
	realRouter, ok := route.(*config.Router)
	if ok {
		return realRouter
	}
	return nil
}
