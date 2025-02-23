package main

import (
	"gin-market-api/controllers"
	"gin-market-api/infra"
	"gin-market-api/middlewares"
	"gin-market-api/repositories"
	"gin-market-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	ItemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", ItemController.FindAll)
	itemRouterWithAuth.GET("/:id", ItemController.FindById)
	itemRouterWithAuth.POST("", ItemController.Create)
	itemRouterWithAuth.PUT("/:id", ItemController.Update)
	itemRouterWithAuth.DELETE("/:id", ItemController.Delete)

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run("localhost:8080")
}
