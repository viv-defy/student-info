package main

import (
	"fmt"
	"student-info/config"
	"student-info/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	router := gin.Default()

	router.GET("/healthcheck", handlers.HealthCheck)
	v1 := router.Group("/v1")
	{
		v1.GET("/students", handlers.GetAllStudents)
		v1.GET("/students/:id", handlers.GetStudentByID)
		v1.POST("/students", handlers.CreateStudent)
		v1.PUT("/students/:id", handlers.UpdateStudent)
		v1.DELETE("/students/:id", handlers.DeleteStudent)
	}

	route := fmt.Sprintf(":%v", config.Port)
	router.Run(route)
}
