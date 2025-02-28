package main

import (
	"gin-market-api/controllers"
	"gin-market-api/infra"
	"gin-market-api/middlewares"
	"gin-market-api/repositories"
	"gin-market-api/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	ItemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
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

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	r := setupRouter(db)

	r.Run("localhost:8080")
}
