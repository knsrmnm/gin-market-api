package main

import (
	"gin-market-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "product1", Price: 1000, Description: "sample data", SoldOut: false},
		{ID: 2, Name: "product2", Price: 2000, Description: "sample data", SoldOut: true},
		{ID: 3, Name: "product3", Price: 3000, Description: "sample data", SoldOut: false},
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080")
}
