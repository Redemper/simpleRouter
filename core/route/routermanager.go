package route

import (
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"simpleRouter/core/filter"
	"strings"
	"sync"
)

var routeLoadType = flag.String("route-load-type", "db", "the way to load routers,Optional value:db,yaml")

var uriRouterMap sync.Map
var nameRouterMap sync.Map

var delegateCache sync.Map

func AddRouter(r *Router) {
	uriRouterMap.Store(r.OriginUri, r)
	nameRouterMap.Store(r.Name, r)
	filters := r.Filters
	split := strings.Split(filters, ",")
	var ds []*filter.Delegate
	for _, s := range split {
		delegate := filter.GetDelegateByName(s)
		ds = append(ds, delegate)
	}
	trip := r.RTrip()
	d := new(filter.Delegate)
	d.Fn = trip
	d.Name = r.Name + "trip"
	ds = append(ds, d)
	delegateCache.Store(r.Name, ds)
}

func FilterRequest(context *gin.Context, uri string) error {
	load, ok := uriRouterMap.Load(uri)
	if !ok {
		return errors.New("cant find router by uri")
	}
	router := load.(*Router)
	delegates, ok := delegateCache.Load(router.Name)
	if !ok {
		return errors.New("cant find filter by uri")
	}
	ds := delegates.([]*filter.Delegate)
	for _, d := range ds {
		d.Fn(context)
	}
	return nil
}

func init() {
	flag.Parse()
	log.Println("routeLoadType is ", routeLoadType)
	var routers []*Router
	switch *routeLoadType {
	case "yaml":
		routers = InitRouterFromYaml()
	case "db":
		fallthrough
	default:
		routers = InitRouterFromDB()
	}
	log.Println("routeLoad finished,routes :")
	for _, r := range routers {
		AddRouter(r)
	}
}

func GetAllRouter() []*Router {
	var routes []*Router
	uriRouterMap.Range(func(key, value interface{}) bool {
		router, ok := value.(*Router)
		if ok {
			routes = append(routes, router)
		}
		return true
	})
	return routes
}
