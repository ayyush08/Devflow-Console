package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ayyush08/keploy-dashboard/routes"
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

	if PORT == "" {
		PORT = ":8080"
	}

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
