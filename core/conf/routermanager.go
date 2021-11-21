package conf

import (
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"simpleRouter/core/cache"
	"simpleRouter/core/component"
	"simpleRouter/core/filter"
	"strings"
)

var routeLoadType = flag.String("route-load-type", "db", "the way to load routers,Optional value:db,yaml")

const (
	RouterPrefix     = "routersCache_"
	NameRouterPrefix = "nameRoutersCache_"
	DelegatePrefix   = "delegateCache_"
)

func AddRouter(r *component.Router) {
	_ = cache.StoreWithoutContext(RouterPrefix+r.OriginUri, r)
	_ = cache.StoreWithoutContext(NameRouterPrefix+r.Name, r)
	filters := r.Filters
	split := strings.Split(filters, ",")
	var ds []*filter.Delegate
	for _, s := range split {
		delegate := filter.GetDelegateByName(s)
		ds = append(ds, delegate)
	}
	// add router trip to filter tail.
	trip := r.RTrip()
	d := new(filter.Delegate)
	d.Fn = trip
	d.Name = r.Name + "trip"
	ds = append(ds, d)
	_ = cache.StoreWithoutContext(DelegatePrefix+r.Name, ds)
}

/**
core func.start filter request.
step 1: get router by uri
step 2: get filter by router
step 3: fire all filters.
*/
func FilterRequest(context *gin.Context, uri string) error {
	load, err := cache.LoadWithoutContext(uri)
	if err != nil {
		return errors.New("cant find router by uri")
	}
	router := load.(*component.Router)
	delegates, err := cache.LoadWithoutContext(DelegatePrefix + router.Name)
	if err != nil {
		return errors.New("cant find filter by uri")
	}
	ds := delegates.([]*filter.Delegate)
	for _, d := range ds {
		d.Fn(context)
	}
	return nil
}

func initRouter() {
	flag.Parse()
	log.Println("routeLoadType is ", routeLoadType)
	var routers []*component.Router
	switch *routeLoadType {
	case "yaml":
		routers = initRouterFromYaml()
	case "db":
		fallthrough
	default:
		routers = initRouterFromDB()
	}
	log.Println("routeLoad finished,routes :")
	for _, r := range routers {
		AddRouter(r)
	}
}

//func GetAllRouter() []*component.Router {
//    var routes []*component.Router
//    cache.LoadWithoutContext("")
//    uriRouterMap.Range(func(key, value interface{}) bool {
//        router, ok := value.(*component.Router)
//        if ok {
//            routes = append(routes, router)
//        }
//        return true
//    })
//    return routes
//}

func initRouterFromYaml() []*component.Router {
	routers := GetYamlRouters()
	if nil == routers {
		return nil
	}
	return routers.Routers
}

type YamlRouters struct {
	Routers []*component.Router `yaml:"router"`
}

func initRouterFromDB() []*component.Router {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&component.Router{})
	var routes []*component.Router
	db.Where(" enabled = 1 ").Find(&routes)
	return routes
}
