package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ayyush08/keploy-dashboard/config"
	"github.com/ayyush08/keploy-dashboard/db"
	"github.com/ayyush08/keploy-dashboard/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var githubAuth = config.NewGithubAuth()

func Login(c *gin.Context) {
    state := config.GenerateState()
    c.SetCookie("oauth_state", state, 3600, "/", "localhost", false, true)
    url := githubAuth.Config.AuthCodeURL(state)
    fmt.Println("Redirecting to:", url)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func Callback(c *gin.Context) {
	storedState, err := c.Cookie("oauth_state")
	if err != nil || storedState != c.Query("state") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid state"})
		return
	}

	code := c.Query("code")
	token, err := githubAuth.Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	userData, err := githubAuth.GetUserInfo(token)

	fmt.Println("User data:", userData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
		return
	}

	email,ok := userData["email"].(string)

	if !ok {
		email = ""
		fmt.Println("Warning: Email not found or is not a string for github response")
	}

	user := models.User{
		GithubID:    fmt.Sprintf("%v", userData["id"]),
		Username:    userData["login"].(string),
		Email:       email,
		AccessToken: token.AccessToken,
		AvatarURL:   userData["avatar_url"].(string),
		CreatedAt:   time.Now(),
	}

	dbClient := db.GetClient()

	collection  := db.GetCollection(dbClient,"users")
	filter := bson.M{"github_id": user.GithubID}
	update := bson.M{"$set": user}
	_, err = collection.UpdateOne(context.Background(), filter, update, options.UpdateOne().SetUpsert(true))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.SetCookie("user_id", user.GithubID, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "username": user.Username})
}


func GetUser(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}