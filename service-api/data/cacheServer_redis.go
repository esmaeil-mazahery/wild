package data

import (
	"context"
	"time"

	"github.com/EsmaeilMazahery/wild/caching"
	"github.com/go-redis/redis/v8"
)

//RedisCacheStore ...
type RedisCacheStore struct {
	cacheServer *caching.RedisCacheServer
}

//NewRedisCacheStore ...
func NewRedisCacheStore() *RedisCacheStore {
	return &RedisCacheStore{
		cacheServer: caching.GetInstanceCacheServer(),
	}
}

//Set Set Key Value To Cache Collection
func (store *RedisCacheStore) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, store.cacheServer.GetTimeout())
	defer cancel()

	err := store.cacheServer.GetClient().Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

// Get finds a member by ID
func (store *RedisCacheStore) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, store.cacheServer.GetTimeout())
	defer cancel()

	val, err := store.cacheServer.GetClient().Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		panic(err)
	}

	return val, nil
}

// Del finds a member by ID
func (store *RedisCacheStore) Del(ctx context.Context, key string) error {
	ctx, cancel := context.WithTimeout(ctx, store.cacheServer.GetTimeout())
	defer cancel()

	err := store.cacheServer.GetClient().Del(ctx, key).Err()
	if err != nil {
		panic(err)
	}

	return nil

}
