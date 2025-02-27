package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ayyush08/keploy-dashboard/db"
	"github.com/ayyush08/keploy-dashboard/handlers"
	"github.com/ayyush08/keploy-dashboard/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
        log.Println("Warning: Could not load .env file:", err)
    }

	_, dbErr := db.ConnectDB()

	if dbErr != nil {
		log.Fatal("Error connecting to MongoDB: ", dbErr)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	var PORT = os.Getenv("PORT")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running 😁!!",
		})
	})


	r.GET("/login",handlers.Login)
	r.GET("/api/auth/callback/github",handlers.Callback)


	protected := r.Group("/api").Use(middlewares.AuthMiddleware())
	{
		protected.GET("/user",handlers.GetUser)
	}



	ginErr := r.Run(PORT)

	if ginErr != nil {
		log.Fatal("Error starting gin: ",ginErr)

	}

}
