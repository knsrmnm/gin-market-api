package main

import (
	"gin-market-api/infra"
	"gin-market-api/models"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}
	code := m.Run()
	os.Exit(code)
}

func setupTestData(db *gorm.DB) {
	items := []models.Item{
		{Name: "item1", Price: 100, Description: "item1 description", SoldOut: false, UserId: 1},
		{Name: "item2", Price: 200, Description: "item2 description", SoldOut: true, UserId: 1},
		{Name: "item3", Price: 300, Description: "item3 description", SoldOut: false, UserId: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "password1"},
		{Email: "test2@example.com", Password: "password2"},
	}

	for _, user := range users {
		db.Create(&user)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.Item{}, &models.User{})
	setupTestData(db)

	router := setupRouter(db)
	return router
}
