package middlewares

import (
	"net/http"

	"github.com/ayyush08/keploy-dashboard/db"
	"github.com/ayyush08/keploy-dashboard/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func AuthMiddleware() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		userID, err := c.Cookie("user_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		dbClient := db.GetClient()
		collection := db.GetCollection(dbClient, "users")
		var user models.User
		err = collection.FindOne(c, bson.M{"github_id": userID}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user) // Store user in context for later use
		c.Next()
	}
}
