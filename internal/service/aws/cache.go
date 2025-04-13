package aws

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisCache(host string, port int, password string, db int) *RedisCache {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password, // empty string if no password
		DB:       db,       // 0 is default DB
	})

	return &RedisCache{
		Client: client,
		Ctx:    ctx,
	}
}

func (r *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, ttl).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisCache) SetJSON(key string, val any, ttl time.Duration) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return r.Set(key, string(b), ttl)
}

func (r *RedisCache) GetJSON(key string, target any) error {
	data, err := r.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), target)
}
