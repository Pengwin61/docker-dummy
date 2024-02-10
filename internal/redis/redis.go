package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ClientRedis *redis.Client

func InitRedisClient() {

	ctx := context.Background()

	addr := "localhost:6379"

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Printf("connection to redis:%s\nping:%s\nerr:%s\n", addr, pong, err)

	ClientRedis = client
}
