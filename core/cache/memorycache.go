package cache

import (
	"context"
	"errors"
	"sync"
)

type MemoryCache struct {
	c  interface{}
	mt McType
}

var CtxDone = errors.New("context is done")
var NoneCache = errors.New("no cache type")

type McType int

const (
	SYCNMAP McType = iota
	MAP
)

func NewMemoryCache(mt McType) *MemoryCache {
	m := new(MemoryCache)
	switch mt {
	case SYCNMAP:
		m.c = &sync.Map{}
	case MAP:
		m.c = make(map[string]interface{})
	default:
		m.c = make(map[string]interface{})
	}
	m.mt = mt
	return m
}

func (mc *MemoryCache) Load(ctx context.Context, key string) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, CtxDone
	default:

	}
	switch mc.c.(type) {
	case sync.Map:
		m := mc.c.(sync.Map)
		return m.Load(key), nil
	case map[string]interface{}:
		m := mc.c.(map[string]interface{})
		return m[key], nil
	}
	return nil, NoneCache
}
func (mc *MemoryCache) Store(ctx context.Context, key string, value interface{}) error {
	select {
	case <-ctx.Done():
		return CtxDone
	default:

	}
	switch mc.c.(type) {
	case sync.Map:
		m := mc.c.(sync.Map)
		m.Store(key, value)
	case map[string]interface{}:
		m := mc.c.(map[string]interface{})
		m[key] = value
	}
	return nil
}
func (mc *MemoryCache) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return CtxDone
	default:

	}
	switch mc.c.(type) {
	case sync.Map:
		m := mc.c.(sync.Map)
		m.Delete(key)
	case map[string]interface{}:
		m := mc.c.(map[string]interface{})
		m[key] = nil
	}
	return nil
}
