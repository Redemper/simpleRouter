package filter

import (
	"github.com/gin-gonic/gin"
	"simpleRouter/core/route/config"
	"sync"
)

//type Filter interface {
//	Filter(context *gin.Context) *http.Response
//}
//
type Ordered interface {
	GetOrder() int
}

func FilterRequest(context *gin.Context, uri string) {
	delegate := getDelegate(uri)
	delegate.fc.Apply(context)
}

//
//type OrderedFilter interface {
//    Filter
//    Ordered
//}

var delegateMap sync.Map

type FilterChan interface {
	Apply(context *gin.Context)
	NextFilter() FilterChan
}

type Delegate struct {
	fc  FilterChan
	uri string
}

func getDelegate(uri string) *Delegate {
	delegate, ok := delegateMap.Load(uri)
	if ok {
		return delegate.(*Delegate)
	}
	d := new(Delegate)
	d.fc = findFiltersByUri(uri)
	d.uri = uri
	delegateMap.Store(uri, d)
	return d
}

func findFiltersByUri(uri string) FilterChan {
	//result := make([]FilterChan,10)
	//result = append(result, new(timeWatchFilter))
	//return new(timeWatchFilter)
	r := new(config.Router)
	r.Name = "test"
	r.TargetUri = "http://localhost:2258"
	return initRouterFilter(r)
}
