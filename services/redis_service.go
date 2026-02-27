package services

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	ctx := context.Background()

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" { redisUrl = "localhost:6379" }

	rdb = redis.NewClient(&redis.Options{
		Addr: redisUrl,
		Password: "",
		DB: 0,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil { panic(err) }
}
