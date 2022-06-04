package cache

import (
	"encoding/json"
	"github.com/abaron10/Posts-API-Golang/model"
	"github.com/go-redis/redis/v7"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *model.Post) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *model.Post {
	client := cache.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	post := model.Post{}
	json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	return &post
}
