package main

import (
	"gin-market-api/infra"
	"gin-market-api/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate")
	}
}
