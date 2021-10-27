package filter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
)

// delegatemap.store and load delegate.
var delegateMap sync.Map

// errors
var DuplicateNameErr = errors.New("already has filter with this name")

// filter delegate.
type Delegate struct {
	Fn   func(context *gin.Context)
	Name string
}

func GetDelegateByName(name string) *Delegate {
	d, ok := delegateMap.Load(name)
	if ok {
		delegate := d.(*Delegate)
		return delegate
	}
	return nil
}

func PutDelegate(d *Delegate) error {
	_, ok := delegateMap.Load(d.Name)
	if ok {
		return DuplicateNameErr
	}
	delegateMap.Store(d.Name, d)
	return nil
}
