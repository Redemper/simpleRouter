package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

//type Filter interface {
//	Filter(context *gin.Context) *http.Response
//}
//
type Ordered interface {
	GetOrder() int
}

//
//type OrderedFilter interface {
//    Filter
//    Ordered
//}

var delegateMap sync.Map

type FilterChan interface {
	Apply(context *gin.Context) *http.Response
	NextFilter() FilterChan
}

type Delegate struct {
	Fc  FilterChan
	Uri string
}

func (d *Delegate) doFilter(context *gin.Context) {
	d.Fc.Apply(context)
}

func GetDelegate(uri string) *Delegate {
	delegate, ok := delegateMap.Load(uri)
	if ok {
		return delegate.(*Delegate)
	}
	d := new(Delegate)
	d.Fc = findFiltersByUri(uri)
	d.Uri = uri
	delegateMap.Store(uri, d)
	return d
}

func findFiltersByUri(uri string) FilterChan {
	//result := make([]FilterChan,10)
	//result = append(result, new(TimeWatchFilter))
	return new(TimeWatchFilter)
}
