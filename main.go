package main

import (
	"log"
	"net/http"
	"petshop/config"
	"petshop/controller"
	"petshop/repository"
	"petshop/router"
	"petshop/service"
	"petshop/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.DatabaseConnection()
	validate := validator.New()

	// Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	// Services
	userService, err := service.NewUserServiceImpl(userRepository, validate)
	if err != nil {
		log.Fatalf("Error initializing User service: %v", err)
	}
	authService := service.NewAuthService(userRepository)

	// Controllers
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)

	// Router
	routes := router.UserRouter(userController, authController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	utils.ErrorPanic(err)
}
