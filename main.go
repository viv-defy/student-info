package main

import (
	"fmt"
	"student-info/config"
	"student-info/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	router := gin.Default()

	router.GET("/healthcheck", handlers.HealthCheck)

	route := fmt.Sprintf(":%v", config.Port)
	router.Run(route)
}
