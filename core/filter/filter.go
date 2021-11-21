package filter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"simpleRouter/core/cache"
)

const (
	DelegateMapPrefix = "delegateMapPrefix_"
)

// errors
var DuplicateNameErr = errors.New("already has filter with this name")

// filter delegate.
type Delegate struct {
	Fn   func(context *gin.Context)
	Name string
}

func GetDelegateByName(name string) *Delegate {
	d, err := cache.LoadWithoutContext(DelegateMapPrefix + name)
	if err == nil {
		delegate := d.(*Delegate)
		return delegate
	}
	return nil
}

func PutDelegate(d *Delegate) error {
	v, err := cache.LoadWithoutContext(d.Name)
	//_, ok := delegateMap.Load(d.Name)
	if err != nil {
		return err
	} else if v != nil {
		return DuplicateNameErr
	}
	cache.StoreWithoutContext(d.Name, d)
	return nil
}
