package filter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
	"simpleRouter/core/route/config"
	"sync"
)

type Ordered interface {
	GetOrder() int
}

func FilterRequest(context *gin.Context, uri string) {
	delegate := getDelegate(uri)
	delegate.fc.Apply(context)
}

var delegateMap sync.Map

var fcMap sync.Map

var WrongTypeErr = errors.New("this is not Filter.please check")
var NotNameErr = errors.New("this filter has't name")
var DuplicateNameErr = errors.New("already has filter with this name")

type Filter interface {
	Apply(context *gin.Context)
	Name() string
}

type OrderFilter interface {
	Ordered
	Filter
}

type Delegate struct {
	fc  Filter
	uri string
}

func getDelegate(uri string) *Delegate {
	reflect.TypeOf("")
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

func findFiltersByUri(uri string) Filter {
	//result := make([]Filter,10)
	//result = append(result, new(timeWatchFilter))
	//return new(timeWatchFilter)

	r := new(config.Router)
	r.Name = "test"
	r.TargetUri = "http://localhost:2258"
	return initRouterFilter(r)
}

// add filter to filter map
func AddFilter(i interface{}) error {
	filter, ok := i.(Filter)
	if !ok {
		return WrongTypeErr
	}
	name := filter.Name()
	if len(name) == 0 {
		return NotNameErr
	}
	_, loaded := fcMap.Load(name)
	if loaded {
		// alread exist
		return DuplicateNameErr
	}
	fcMap.Store(filter.Name(), filter)
	return nil
}
