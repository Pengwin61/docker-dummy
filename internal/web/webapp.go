package web

import (
	"docker-dummy/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WebParams struct {
	redisHost   string
	rabbitHost  string
	rabbitPort  string
	rabbitQueue string
}

var webParams *WebParams

func InitGun(config *config.Config) {
	webParams = getParams(config)

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

func getParams(config *config.Config) *WebParams {

	return &WebParams{
		redisHost:   config.Redis.Host,
		rabbitHost:  config.RabbitMQ.Host,
		rabbitPort:  config.RabbitMQ.Port,
		rabbitQueue: config.RabbitMQ.QueueName}
}
