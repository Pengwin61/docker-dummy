package main

import (
	"docker-dummy/internal/config"
	"docker-dummy/internal/connections"
	"docker-dummy/internal/redis"
	"docker-dummy/internal/web"
	"fmt"
	"time"
)

func main() {
	conf := config.InitConfig()

	if conf.Server {
		startServer()
	} else if conf.Client {
		startClient()
	} else {
		fmt.Println("Please specify server or client")
	}
}

func startServer() {
	fmt.Println("Starting server....")

	connections.InitAllConnections(redis.RedisIp, "", 0)

	defer connections.Con.Redis.Close()

	go web.InitGun()

	web.ReadInQueue()
}

func startClient() {
	var Count = 1000000

	fmt.Println("Starting client....")

	connections.InitAllConnections(redis.RedisIp, "", 0)

	defer connections.Con.Redis.Close()

	for {

		connections.Con.Redis.Set("name", "Elliot")
		connections.Con.Redis.Set("surname", "Alderson")
		connections.Con.Redis.Set("msg", "Hello, Friend!")

		for j := 0; j < Count; j++ {

			time.Sleep(1 * time.Second)

			connections.Con.RabbitMQ.Send(fmt.Sprint("MR.ROBOT", ":", j))
		}

		time.Sleep(20 * time.Second)
		fmt.Println("Sleeep...")

	}
}
