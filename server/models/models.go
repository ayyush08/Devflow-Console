package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
    ID          bson.ObjectID `bson:"_id,omitempty"`
    GithubID    string             `bson:"github_id"`
    Username    string             `bson:"username"`
    Email       string             `bson:"email,omitempty"`
    AccessToken string             `bson:"access_token"`
    AvatarURL   string             `bson:"avatar_url,omitempty"`
    Repositories []string          `bson:"repositories"`
    CreatedAt   time.Time         `bson:"created_at"`
}