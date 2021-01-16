package data

import (
	"context"
	"time"
)

// CacheServer is ...
type CacheServer interface {
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
