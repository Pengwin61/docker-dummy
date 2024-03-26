package web

import (
	"docker-dummy/internal/core"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// func getResponseHTML(c *gin.Context) {
// 	c.HTML(http.StatusOK, "index.html", gin.H{

// 		"time":     time.Now().Format("2006-01-02 15:04:05"),
// 		"hostname": core.GetHostname(),
// 		"ip":       core.GetIP(),
// 		"version":  webParams.appVersion,

// 		"redisHost":   webParams.redisHost,
// 		"redisStatus": checkRedis(),
// 		"redisResponse": RedisResponse{
// 			Name:    getInRedis("name"),
// 			Surname: getInRedis("surname"),
// 		},

// 		"rabbitHost":  webParams.rabbitHost,
// 		"rabbitPort":  webParams.rabbitPort,
// 		"rabbitQueue": webParams.rabbitQueue,
// 		"rabbitResponse": RabbitResponse{
// 			Msg: msg,
// 		},

// 		"dbHost":     webParams.databaseHost,
// 		"dbName":     webParams.databaseName,
// 		"dbStatus":   "TO_DO",
// 		"dbResponse": "TO_DO",
// 	})
// }

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
				Hostname: webParams.databaseHost,
				DbName:   webParams.databaseName,
				Status:   checkDb(),
				DbResponse: DbResponse{
					Name:    "TO_DO",
					Surname: "TO_DO",
				},
			},
		},
	}
	c.JSON(http.StatusOK, response)
}
