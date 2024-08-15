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

	"github.com/go-playground/validator/v10"
)

func main() {
	db := postgres.OpenConnection()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validator.New())
	userHandler := handler.NewUserHandler(userService)
	userRoutes := route.NewUserRoute(userHandler)
	server := api.NewServer(userRoutes)

	server.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
