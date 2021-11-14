package cache

import (
    "context"
)

type ICache interface {
    Load(ctx context.Context,key string) (interface{} ,error)
    Store(ctx context.Context,key string, value interface{}) error
    Delete(ctx context.Context,key string) error
}
