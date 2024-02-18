package web

import (
	"docker-dummy/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

var redisHost string
var rabbitHost string
var rabbitPort string

func InitGun(config *config.Config) {

	redisHost = config.Redis.Host
	rabbitHost = config.RabbitMQ.Host
	rabbitPort = config.RabbitMQ.Port

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	RunApi(r)

	err := r.Run(fmt.Sprint(":", "80"))
	if err != nil {
		panic(err)
	}

}

func RunApi(r *gin.Engine) {

	r.GET("/", getResponse)
}
