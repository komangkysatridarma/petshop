package router

import (
	"petshop/controller"
	"petshop/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(userController *controller.UserController, authController *controller.AuthController) *gin.Engine {
	r := gin.Default()

	r.POST("/login", authController.Login)

	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("", userController.FindAll)
		userRoutes.GET("/:id", userController.FindById)
		userRoutes.POST("", userController.Create)
		userRoutes.PATCH("/:id", userController.Update)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	return r
}
