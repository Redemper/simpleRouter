package router

import (
	"bufio"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"simpleRouter/core/cache"
	"simpleRouter/core/discover"
	"simpleRouter/core/filter"
	"simpleRouter/core/pojo"
	"strings"
)

var routeLoadType = flag.String("route-load-type", "db", "the way to load routers,Optional value:db,yaml")

const (
	RouterPrefix     = "routersCache_"
	NameRouterPrefix = "nameRoutersCache_"
	DelegatePrefix   = "delegateCache_"
)

func AddRouter(r *pojo.Router) {
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
	trip := trip(r)
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
	router := load.(*pojo.Router)
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
	var routers []*pojo.Router
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

//func GetAllRouter() []*pojo.Router {
//   var routes []*pojo.Router
//   cache.LoadWithoutContext("")
//   uriRouterMap.Range(func(key, value interface{}) bool {
//       router, ok := value.(*pojo.Router)
//       if ok {
//           routes = append(routes, router)
//       }
//       return true
//   })
//   return routes
//}

func trip(r *pojo.Router) func(context *gin.Context) {
	return func(context *gin.Context) {
		uri := discover.GetTagetUriByOriginUri(r.OriginUri)
		req := context.Request
		proxy, err := url.Parse(uri)
		if err != nil {
			log.Printf("error in parse addr: %v", err)
			context.String(500, "error")
			return
		}
		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)
		if err != nil {
			log.Printf("error in roundtrip: %v", err)
			context.String(500, "error")
			return
		}
		for k, vv := range resp.Header {
			for _, v := range vv {
				context.Header(k, v)
			}
		}
		defer resp.Body.Close()
		bufio.NewReader(resp.Body).WriteTo(context.Writer)
	}
}
