package main

import (
	"fmt"
	"log"
	"net/http"
	"raihan-athallah-api-golang/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Membuat router Gin
	router := gin.Default()

	// Configure CORS
	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders:     []string{"Content-Type", "Authorization", "token"}, // Add the "token" header here
				AllowCredentials: true,
			},
		),
	)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := "8080"
	// Menjalankan server
	fmt.Printf("Server is running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	//Initialize redis connection
	config.InitRedis()

	// Initialize the database connection
	config.InitDB()

	// Migrate the database schema
	// config.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
}
