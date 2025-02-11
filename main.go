package main

import (
	"gin-market-api/controllers"
	"gin-market-api/infra"
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

	r := gin.Default()
	r.GET("/items", ItemController.FindAll)
	r.GET("/items/:id", ItemController.FindById)
	r.POST("/items", ItemController.Create)
	r.PUT("/items/:id", ItemController.Update)
	r.DELETE("/items/:id", ItemController.Delete)
	r.Run("localhost:8080")
}
