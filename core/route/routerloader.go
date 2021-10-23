package route

import (
	"flag"
	"log"
	"simpleRouter/core/route/config"
)

var routeLoadType = flag.String("route-load-type", "yaml", "the way to load routers,Optional value:db,yaml")

func init() {
	flag.Parse()
	log.Println("routeLoadType is ", routeLoadType)
	switch *routeLoadType {
	case "yaml":
		config.InitRouterFromYaml()
	case "db":
		fallthrough
	default:
		config.InitRouterFromDB()
	}
	log.Println("routeLoad finished,routes :")
	for _, r := range config.GetAllRouter() {
		log.Println(*r)
	}
}
