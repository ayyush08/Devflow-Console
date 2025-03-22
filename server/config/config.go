package config

import (
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

func GetGithubToken() string {
	return os.Getenv("GITHUB_ACCESS_TOKEN")
}

type MetricsCacheItem struct {
    Type  string
    Value interface{}
}


var GlobalCache = cache.New(10*time.Minute,15*time.Minute)
