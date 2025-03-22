package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ayyush08/devflow-console/middlewares"
	"github.com/ayyush08/devflow-console/routes"
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

	FRONTEND_URL, exists := os.LookupEnv("FRONTEND_URL")
	if !exists {
		log.Println("FRONTEND_URL not found, using default")
		FRONTEND_URL = "http://localhost:3000"
	}

	r.Use(middlewares.CorsMiddleware(FRONTEND_URL))
	var PORT = os.Getenv("PORT")

	if PORT == "" {
		PORT = ":8080"
		log.Println("PORT not found in .env, using default:", PORT)
	}


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
