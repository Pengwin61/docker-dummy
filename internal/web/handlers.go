package web

import (
	"docker-dummy/internal/connections"
	"docker-dummy/internal/core"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getResponseHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{

		"time":     time.Now().Format("2006-01-02 15:04:05"),
		"hostname": core.GetHostname(),
		"ip":       core.GetIP(),
		"version":  webParams.appVersion,

		"redisHost":   webParams.redisHost,
		"redisStatus": checkRedis(),
		"redisResponse": RedisResponse{
			Name:    getInRedis("name"),
			Surname: getInRedis("surname"),
		},

		"rabbitHost":  webParams.rabbitHost,
		"rabbitPort":  webParams.rabbitPort,
		"rabbitQueue": webParams.rabbitQueue,
		"rabbitResponse": RabbitResponse{
			Msg: msg,
		},

		"dbHost":     "TO_DO",
		"dbName":     "TO_DO",
		"dbStatus":   "TO_DO",
		"dbResponse": "TO_DO",
	})
}

func getResponseJSON(c *gin.Context) {
	response := response{
		TimeStamp: time.Now().Format("2006-01-02 15:04:05"),
		System: System{
			Hostname: core.GetHostname(),
			Ip:       core.GetIP(),
			Version:  webParams.appVersion,
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
				Status:   "TO_DO",
				RabbitResponse: RabbitResponse{
					Msg: msg,
				},
			},
			Database: Database{
				Hostname: "TO_DO",
				Status:   "TO_DO",
				DbResponse: DbResponse{
					Name:    "TO_DO",
					Surname: "TO_DO",
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
