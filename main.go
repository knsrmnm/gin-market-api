package main

import (
	"gin-market-api/controllers"
	"gin-market-api/models"
	"gin-market-api/repositories"
	"gin-market-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "product1", Price: 1000, Description: "sample data", SoldOut: false},
		{ID: 2, Name: "product2", Price: 2000, Description: "sample data", SoldOut: true},
		{ID: 3, Name: "product3", Price: 3000, Description: "sample data", SoldOut: false},
	}

	itemRepository := repositories.NewItemMemoryRepository(items)
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
