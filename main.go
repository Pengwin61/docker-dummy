package main

import (
	"context"
	"docker-dummy/internal/redis"
	"docker-dummy/internal/web"
	"docker-dummy/internal/wrapper"
)

func main() {

	redis.InitRedisClient()

	err := redis.ClientRedis.Set(context.Background(), "name", "Elliot", 0).Err()
	wrapper.Err(err)

	err = redis.ClientRedis.Set(context.Background(), "surname", "Alderson", 0).Err()
	wrapper.Err(err)

	web.InitGun()

}
