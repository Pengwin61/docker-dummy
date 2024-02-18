package main

import (
	"docker-dummy/internal/config"
	"docker-dummy/internal/connections"
	"docker-dummy/internal/web"
	"fmt"
	"log"
	"time"
)

func main() {
	conf := config.GetConfig()

	if conf.App.Server {
		startServer(conf)
	} else if conf.App.Client {
		startClient(conf)
	} else {
		fmt.Println("Please specify server or client")
	}
}

func startServer(config *config.Config) {
	log.Println("Starting server....")
	log.Printf("redis enable: %t, rabbit enable: %t\n", config.Redis.Enable, config.RabbitMQ.Enable)

	connections.InitAllConnections(config)

	if config.Redis.Enable {
		defer connections.Con.Redis.Close()
	}
	if config.RabbitMQ.Enable {
		go web.ReadInQueue()
	}

	web.InitGun(config)
}

func startClient(config *config.Config) {
	var Count = 1000000

	log.Println("Starting client....")
	log.Printf("redis enable: %t, rabbit enable: %t\n", config.Redis.Enable, config.RabbitMQ.Enable)

	connections.InitAllConnections(config)

	defer connections.Con.Redis.Close()
	for {

		if config.Redis.Enable {
			connections.Con.Redis.Set("name", "Elliot")
			connections.Con.Redis.Set("surname", "Alderson")
			connections.Con.Redis.Set("msg", "Hello, Friend!")
		}

		for j := 0; j < Count; j++ {

			time.Sleep(1 * time.Second)

			if config.RabbitMQ.Enable {
				connections.Con.RabbitMQ.Send(fmt.Sprint("fsociety", j, ".dat"))
			} else {
				break
			}

		}

		time.Sleep(20 * time.Second)
		fmt.Println("Sleeep...")

	}

}
