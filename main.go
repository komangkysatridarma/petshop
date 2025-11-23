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
	branchRepository := repository.NewBranchRepositoryImpl(db)
	roleRepository := repository.NewRoleRepositoryImpl(db)
	categoryRepository := repository.NewCategoryRepositoryImpl(db)
	productRepository := repository.NewProductRepositoryImpl(db)

	// Services
	userService, err := service.NewUserServiceImpl(userRepository, validate)
	if err != nil {
		log.Fatalf("Error initializing User service: %v", err)
	}
	authService := service.NewAuthService(userRepository)

	branchService, err := service.NewBranchServiceImpl(branchRepository, validate)
	if err != nil {
		log.Fatalf("Error initializing Branch service: %v", err)
	}

	roleService := service.NewRoleServiceImpl(roleRepository, validate)
	categoryService := service.NewCategoryServiceImpl(categoryRepository, validate)
	productService := service.NewProductServiceImpl(productRepository, validate)

	// Controllers
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(authService)
	branchController := controller.NewBranchController(branchService)
	roleController := controller.NewRoleController(roleService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)

	// Router
	routes := router.SetupRouter(
		authController,
		userController,
		branchController,
		roleController,
		categoryController,
		productController,
	)

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
