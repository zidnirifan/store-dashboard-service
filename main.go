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
	categoryRepository := repository.NewCategoryRepository(db)

	validate := validator.New()
	userService := service.NewUserService(userRepository, validate)
	categoryService := service.NewCategoryService(categoryRepository, validate)

	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	userRoutes := route.NewUserRoute(userHandler)
	categoryRoute := route.NewCategoryRoute(categoryHandler)
	route := api.Route{
		UserRoute:     userRoutes,
		CategoryRoute: categoryRoute,
	}
	server := api.NewServer(&route)

	server.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
