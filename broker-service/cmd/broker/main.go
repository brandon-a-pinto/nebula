package main

import (
	"fmt"

	"github.com/brandon-a-pinto/nebula/broker-service/configs"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/main/web"
)

func main() {
	config := configs.LoadConfig()

	// Web Server
	webserver := web.NewWebServer(":" + config.WebServerPort)
	fmt.Println("Starting web server on port", config.WebServerPort)
	webserver.Start()
}
