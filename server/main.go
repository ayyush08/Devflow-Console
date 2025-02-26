package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ayyush08/keploy-dashboard/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	_, dbErr := db.ConnectDB()

	if dbErr != nil {
		log.Fatal("Error connecting to MongoDB: ", dbErr)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	var PORT = os.Getenv("PORT")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	ginErr := r.Run(PORT)

	if ginErr != nil {
		log.Fatal("Error starting gin: ",ginErr)

	}

}
