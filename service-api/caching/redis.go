package caching

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/EsmaeilMazahery/wild/infrastructure/constant"
	"github.com/go-redis/redis/v8"
)

//RedisCacheServer used for connect to redis server
type RedisCacheServer struct {
	client  *redis.Client
	timeout time.Duration
}

//GetInstanceCacheServer get a singelton instance of redis server
func GetInstanceCacheServer() *RedisCacheServer {
	return singleton
}

var singleton *RedisCacheServer

func init() {
	timeout, err := strconv.Atoi(os.Getenv("REDIS_TIMEOUT"))
	if err != nil {
		log.Fatalln("REDIS_TIMEOUT is incorrect:", err)
	}

	singleton = &RedisCacheServer{
		client:  getClient(),
		timeout: time.Duration(timeout) * time.Millisecond,
	}
}

func getClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     constant.RedisURI(),
		Password: constant.RedisPassword(),
		DB:       0, // use default DB
	})

	return rdb
}

//GetClient ...
func (server *RedisCacheServer) GetClient() *redis.Client {
	return server.client
}

//GetTimeout ..
func (server *RedisCacheServer) GetTimeout() time.Duration {
	return server.timeout
}
