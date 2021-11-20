package cache

import (
	"context"
	"flag"
)

type ICache interface {
	Load(ctx context.Context, key string) (interface{}, error)
	Store(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

var CacheType = flag.String("choose your caceh type", "memory", "")

func GetCache() ICache {
	flag.Parse()
	switch *CacheType {
	case "redis":
	case "memory":
	default:

	}
	return nil
}
