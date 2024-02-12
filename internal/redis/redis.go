package redis

import (
	"context"
	"docker-dummy/internal/wrapper"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisIp = "192.168.140.185:6379"

type ClientRedis struct {
	client *redis.Client
}

func NewRedisClient(url string, password string, db int) *ClientRedis {

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		wrapper.Err(err)
	} else {
		fmt.Printf("connection to redis:%s\nping:%s\nerr:%s\n", url, pong, err)
	}

	return &ClientRedis{client: client}
}

func (cr *ClientRedis) Get(entity string) string {
	entity, err := cr.client.Get(context.Background(), entity).Result()
	if err != nil {
		fmt.Println(err)
	}
	return entity
}

func (cr *ClientRedis) Set(key, value string) {
	err := cr.client.Set(context.Background(), key, value, 0).Err()
	wrapper.Err(err)
}

func (cr *ClientRedis) Check() *redis.StatusCmd {
	return cr.client.Ping(context.Background())
}

func (cr *ClientRedis) Close() {
	cr.client.Close()
}
