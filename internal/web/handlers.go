package web

import (
	"docker-dummy/internal/connections"
	"docker-dummy/internal/core"
	"docker-dummy/internal/database"
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
				Hostname: webParams.redisHost,
				Status:   checkRedis(),
				RedisResponse: RedisResponse{
					Name:    getInRedis("name"),
					Surname: getInRedis("surname"),
				},
			},
			Rabbit: Rabbit{
				Hostname: fmt.Sprint(webParams.rabbitHost, ":", webParams.rabbitPort),
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

func ReadInQueue() {

	if connections.Con.RabbitMQ != nil {
		messages := connections.Con.RabbitMQ.Get(webParams.rabbitQueue)
		forever := make(chan bool)

		go func() {
			for message := range messages {
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

}

func checkRedis() string {

	if connections.Con.Redis != nil {
		return connections.Con.Redis.Check()
	}
	return ""
}

func getInRedis(entity string) string {
	if connections.Con.Redis != nil {
		return connections.Con.Redis.Get(entity)
	}
	return ""
}
