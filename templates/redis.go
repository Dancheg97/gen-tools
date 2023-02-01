package templates

const RedisGo = `package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
}

func Get(conn string) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: conn,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &Redis{Client: rdb}, err
}

func (r *Redis) Put(ctx context.Context, key string, value []byte) error {
	_, err := r.Client.Set(ctx, key, string(value), time.Second).Result()
	return err
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	rez, err := r.Client.Get(ctx, key).Result()
	return []byte(rez), err
}
`
