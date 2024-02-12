package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitGun() {

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
