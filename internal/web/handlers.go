package web

import (
	"docker-dummy/internal/connections"
	"docker-dummy/internal/core"
	"docker-dummy/internal/database"
	"docker-dummy/internal/rabbitmq"
	"docker-dummy/internal/redis"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getResponse(c *gin.Context) {
	response := response{
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		System: System{
			Hostname: core.GetHostname(),
			Ip:       core.GetIP(),
		},
		ExternalService: ExternalService{
			Redis: Redis{
				Hostname: redis.RedisIp,
				Status:   connections.Con.Redis.Check().Val(),
				RedisResponse: RedisResponse{
					Name:    connections.Con.Redis.Get("name"),
					Surname: connections.Con.Redis.Get("surname"),
				},
			},
			Rabbit: Rabbit{
				Hostname: fmt.Sprint(rabbitmq.RabbitIP, ":", rabbitmq.RabbitPort),
				Status:   "to_do",
				RabbitResponse: RabbitResponse{
					Msg: msg,
				},
			},
			Database: Database{
				Hostname: database.PostgresIp,
				Status:   "to_do",
				DbResponse: DbResponse{
					Name:    "to_do",
					Surname: "to_do",
				},
			},
		},
	}
	c.JSON(http.StatusOK, response)
}

var msg []string

func T() {

	messages := connections.Con.RabbitMQ.Get("testQueue")
	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf(" > Received message: %s\n", message.Body)

			if len(msg) < 500 {
				msg = append(msg, string(message.Body))
			} else {
				log.Println("переполнился")
				msg = nil
			}

		}
	}()
	<-forever
}
