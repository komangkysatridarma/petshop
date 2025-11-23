package router

import (
	"petshop/controller"
	"petshop/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	authController *controller.AuthController,
	userController *controller.UserController,
	branchController *controller.BranchController,
	roleController *controller.RoleController,
	categoryController *controller.CategoryController,
	productController *controller.ProductController,
	// purchaseController *controller.PurchaseController,
) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authController.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/users", userController.FindAll)
		api.GET("/users/:id", userController.FindById)
		api.POST("/users", userController.Create)
		api.PATCH("/users/:id", userController.Update)
		api.DELETE("/users/:id", userController.Delete)

		api.GET("/branches", branchController.FindAll)
		api.GET("/branches/:id", branchController.FindById)
		api.POST("/branches", branchController.Create)
		api.PATCH("/branches/:id", branchController.Update)
		api.DELETE("/branches/:id", branchController.Delete)

		api.GET("/roles", roleController.FindAll)
		api.GET("/roles/:id", roleController.FindById)
		api.POST("/roles", roleController.Create)
		api.PATCH("/roles/:id", roleController.Update)
		api.DELETE("/roles/:id", roleController.Delete)

		api.GET("/categories", categoryController.FindAll)
		api.GET("/categories/:id", categoryController.FindById)
		api.POST("/categories", categoryController.Create)
		api.PATCH("/categories/:id", categoryController.Update)
		api.DELETE("/categories/:id", categoryController.Delete)

		api.GET("/products", productController.FindAll)
		api.GET("/products/:id", productController.FindById)
		api.POST("/products", productController.Create)
		api.PATCH("/products/:id", productController.Update)
		api.DELETE("/products/:id", productController.Delete)

		// api.GET("/purchases", purchaseController.FindAll)
		// api.GET("/purchases/:id", purchaseController.FindById)
		// api.POST("/purchases", purchaseController.Create)
		// api.PATCH("/purchases/:id", purchaseController.Update)
		// api.DELETE("/purchases/:id", purchaseController.Delete)
	}

	return r
}
