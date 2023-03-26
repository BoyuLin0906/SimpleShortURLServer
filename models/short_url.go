package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetOriginalURL(short_url string) string {
	rdb := RedisClient(ctx)
	val, err := rdb.Get(ctx, short_url).Result()
	var original_url string = ""

	if err != nil && err != redis.Nil {
		panic(err)
	} else {
		original_url = val
	}

	return original_url
}

func SetShortURL(short_url string, original_url string) {
	rdb := RedisClient(ctx)
	err := rdb.Set(ctx, short_url, original_url, 0).Err()
	if err != nil {
		panic(err)
	}
}
