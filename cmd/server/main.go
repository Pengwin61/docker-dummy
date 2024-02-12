package main

import (
	"docker-dummy/internal/connections"
	"docker-dummy/internal/rabbitmq"
	"docker-dummy/internal/redis"
	"docker-dummy/internal/web"
)

func main() {
	rabbitmq.Init()

	connections.InitAllConnections(redis.RedisIp, "", 0)

	connections.Con.Redis.Set("name", "Elliot")
	connections.Con.Redis.Set("surname", "Alderson")
	connections.Con.Redis.Set("msg", "Hello, Friend!")

	defer connections.Con.Redis.Close()

	web.InitGun()
}
