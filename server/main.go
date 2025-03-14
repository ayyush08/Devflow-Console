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

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Could not load .env file:", err)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	var PORT = os.Getenv("PORT")
	
	var FRONTEND_URL = os.Getenv("FRONTEND_URL")

	if PORT == "" {
		PORT = ":8080"
	}


	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{FRONTEND_URL}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
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
		routes.PRRoutes(api)
		routes.GetResponse(api) // to be removed later
	}

	ginErr := r.Run(PORT)

	if ginErr != nil {
		log.Fatal("Error starting gin: ", ginErr)
	}

}
