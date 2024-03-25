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
	appVersion  string
}

var webParams *WebParams

func InitGun(config *config.Config) {
	webParams = getParams(config)

	r := gin.Default()
	r.LoadHTMLGlob("templates/html/*.html")

	gin.SetMode(gin.ReleaseMode)

	RunApi(r)

	err := r.Run(fmt.Sprint(":", "80"))
	if err != nil {
		panic(err)
	}

}

func RunApi(r *gin.Engine) {

	r.GET("/", getResponseJSON)

	r.GET("/v1", getResponseHTML)
}

func getParams(config *config.Config) *WebParams {

	return &WebParams{
		redisHost:   config.Redis.Host,
		rabbitHost:  config.RabbitMQ.Host,
		rabbitPort:  config.RabbitMQ.Port,
		rabbitQueue: config.RabbitMQ.QueueName,
		appVersion:  config.App.Version}
}
