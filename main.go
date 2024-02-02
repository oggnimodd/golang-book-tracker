package main

import (
	"net/http"
	"os"
	"strconv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oggnimodd/golang-clerk-practice/context"
	"github.com/oggnimodd/golang-clerk-practice/db"
	"github.com/oggnimodd/golang-clerk-practice/docs"
	"github.com/oggnimodd/golang-clerk-practice/routes"
)

func init() {
	context.DB = db.Init()
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Init gin
	r := gin.Default()

	routes.SetupRoutes(r)

	docs.SwaggerInfo.BasePath = "/api/v1"

	// Check if DEV_MODE is set to true
	devMode, _ := strconv.ParseBool(os.Getenv("DEV_MODE"))

	// Books

	if devMode {

		r.GET("/CLERK_PUBLISHABLE_KEY", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"CLERK_PUBLISHABLE_KEY": os.Getenv("CLERK_PUBLISHABLE_KEY"),
			})
		})
		r.Static("/docs", "./swagger")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	r.Run(":" + port)
}
