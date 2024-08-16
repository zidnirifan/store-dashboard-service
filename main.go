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
	"store-dashboard-service/storage"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := postgres.OpenConnection()
	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	productRepository := repository.NewProductRepository(db)

	cloudinary := storage.NewCloudinary()
	validate := validator.New()
	userService := service.NewUserService(userRepository, validate)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	productService := service.NewProductService(service.ProductServiceParams{
		Repository:         productRepository,
		CategoryRepository: categoryRepository,
		Validate:           validate,
		Cloudinary:         cloudinary,
	})

	userHandler := handler.NewUserHandler(userService)
	categoryHandler := handler.NewCategoryHandler(categoryService)
	productHandler := handler.NewProductHandler(productService)

	userRoutes := route.NewUserRoute(userHandler)
	categoryRoute := route.NewCategoryRoute(categoryHandler)
	productRoute := route.NewProductRoute(productHandler)
	route := api.Route{
		UserRoute:     userRoutes,
		CategoryRoute: categoryRoute,
		ProductRoute:  productRoute,
	}
	server := api.NewServer(&route)

	server.Listen(fmt.Sprintf(":%d", config.GetConfig().Port))
}
