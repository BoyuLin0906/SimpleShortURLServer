package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const host string = "localhost:6379"
const pwd string = ""
const db int = 0

func RedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return client
}
