package router

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"student-info/config"
	"student-info/handlers"
	"time"

	"github.com/gin-gonic/gin"
)

func formatLogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func GetRouter() (*gin.Engine, error) {
	router := gin.New()

	f := getLogFile()
	gin.DefaultWriter = io.MultiWriter(f)

	router.Use(gin.LoggerWithFormatter(formatLogs))
	router.Use(gin.Recovery())

	router.GET("/healthcheck", handlers.HealthCheck)
	v1 := router.Group("/v1")
	{
		v1.GET("/students", handlers.GetAllStudents)
		v1.GET("/students/:id", handlers.GetStudentByID)
		v1.POST("/students", handlers.CreateStudent)
		v1.PUT("/students/:id", handlers.UpdateStudent)
		v1.DELETE("/students/:id", handlers.DeleteStudent)
	}

	return router, nil
}

func getLogFile() *os.File {
	logsPath := filepath.Join(config.RootDir, "student-info.log")
	f, err := os.OpenFile(logsPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error creating log file")
	}

	return f
}
