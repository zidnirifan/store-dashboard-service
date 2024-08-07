package main

import (
	"fmt"
	"store-dashboard-service/api"
	"store-dashboard-service/config"
)

func main() {
	server := api.NewServer()

	server.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
