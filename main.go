package main

import (
	"fmt"
	"student-info/config"
	"student-info/router"
)

func main() {
	config.Init()
	router, err := router.GetRouter()
	if err != nil {
		return
	}

	route := fmt.Sprintf(":%v", config.Port)
	router.Run(route)
}
