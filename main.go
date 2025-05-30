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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your JWT token.

package main

import (
	"log"
	"os"
	"service-pace11/config"
	"service-pace11/routes"
	"service-pace11/wire"

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
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.ConnectDB()
	app := wire.InitApp()
	routes.SetupRoutes(r, app)
	r.Run(os.Getenv("APP_PORT"))
}
