package main

import (
	"fmt"
	"store-dashboard-service/api"
	"store-dashboard-service/api/handler"
	"store-dashboard-service/api/route"
	"store-dashboard-service/config"
	"store-dashboard-service/db/postgres"
	"store-dashboard-service/repository"
	"store-dashboard-service/service"
)

func main() {
	db := postgres.OpenConnection()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userRoutes := route.NewUserRoute(*userHandler)
	server := api.NewServer(userRoutes)

	server.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
