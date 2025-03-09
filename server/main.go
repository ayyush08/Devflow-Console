package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	var PORT = os.Getenv("PORT")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running 😁!!",
		})
	})

	ginErr := r.Run(PORT)

	if ginErr != nil {
		log.Fatal("Error starting gin: ", ginErr)
	}

}
