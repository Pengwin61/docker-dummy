package main

import (
	"docker-dummy/internal/connections"
	"docker-dummy/internal/redis"
	"docker-dummy/internal/web"
	"fmt"
)

func main() {

	fmt.Println("Starting server....")

	connections.InitAllConnections(redis.RedisIp, "", 0)

	defer connections.Con.Redis.Close()

	go web.InitGun()

	web.T()

}
