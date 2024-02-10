package web

import (
	"context"
	"docker-dummy/internal/core"
	"docker-dummy/internal/redis"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitGun() {

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	// RunHandlers(r)
	RunApi(r)

	err := r.Run(fmt.Sprint(":", "80"))
	if err != nil {
		panic(err)
	}
}

func RunApi(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "hey",
			"status":   http.StatusOK,
			"hostname": core.GetHostname(),
			"ip":       core.GetIP(),
			"name":     getEntity("name"),
			"surname":  getEntity("surname")},
		)
	})

}

func getEntity(entity string) string {
	entity, err := redis.ClientRedis.Get(context.Background(), entity).Result()
	if err != nil {
		fmt.Println(err)
	}
	return entity
}
