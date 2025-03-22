package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ayyush08/devflow-console/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	log.SetOutput(os.Stdout)
	gin.DefaultWriter = os.Stdout
	gin.DefaultErrorWriter = os.Stderr

	log.Println("Starting server setup...")
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file:", err)
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	var PORT = os.Getenv("PORT")

	var FRONTEND_URL = os.Getenv("FRONTEND_URL")

	if PORT == "" {
		PORT = ":8080"
		log.Println("PORT not found in .env, using default:", PORT)
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FRONTEND_URL}, // Ensure it's correctly set
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running 😁!!",
		})
	})

	api := r.Group("/api/v1")
	{
		routes.MetricRoutes(api)
	}

	log.Println("Starting server on", PORT)

	if err := r.Run(PORT); err != nil {
		log.Fatal("Error starting Gin server: ", err)
	}

}
