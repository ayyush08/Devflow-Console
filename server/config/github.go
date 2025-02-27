package config

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GithubAuth struct {
	Config *oauth2.Config
}

func NewGithubAuth() *GithubAuth {

	godotenv.Load()
	clientID := os.Getenv("GITHUB_CLIENT_ID")
	clientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	redirectURL := os.Getenv("GITHUB_CALLBACK_URL")

	fmt.Printf("ClientID: %s, ClientSecret: %s, RedirectURL: %s\n", clientID, clientSecret, redirectURL)

	if clientID == "" || clientSecret == "" || redirectURL == "" {
		panic("Missing required GitHub OAuth environment variables")
	}

	return &GithubAuth{
		Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"user:email", "repo"},
			Endpoint:     github.Endpoint,
		},
	}
}

func GenerateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (ga *GithubAuth) GetUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := ga.Config.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		return nil, err
	}
	return userData, nil
}
