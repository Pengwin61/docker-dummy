package main

import (
	"docker-dummy/internal/connections"
	"docker-dummy/internal/redis"
	"fmt"
	"time"
)

var Count = 1000000

func main() {

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
