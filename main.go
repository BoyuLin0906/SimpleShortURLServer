package main

import (
	"log"
	"simple_short_url_server/router"
)

const port string = "8888"

func main() {
	router_main := router.InitServer()
	err := router_main.Run(":" + port)
	if err != nil {
		log.Fatalf("Start server: %+v", err)
	}
}
