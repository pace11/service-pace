// @title Pace Service API
// @version 1.0
// @description This is the API documentation for Pace Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Pace Support
// @contact.url https://www.linkedin.com/in/muhammad-iriansyah-putra-pratama-a0120514b/
// @contact.email rppratama1771@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host service.pace11.my.id
// @BasePath /api

package main

import (
	"log"
	"os"
	"service-pace11/config"
	"service-pace11/routes"

	_ "service-pace11/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	r := gin.Default()
	r.Use(cors.Default())
	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(os.Getenv("APP_PORT"))
}
