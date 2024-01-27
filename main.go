package main

import (
	"fmt"
	"student-info/config"
	"student-info/router"
)

func main() {
	config.Init()
	router := router.GetRouter()

	route := fmt.Sprintf(":%v", config.Port)
	router.Run(route)
}
