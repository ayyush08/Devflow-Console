package config

import (
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

func GetGithubToken() string {
	return os.Getenv("GITHUB_ACCESS_TOKEN")
}


var MetricsCache = cache.New(10*time.Minute,15*time.Minute)

var GeneralMetricsCache = cache.New(10*time.Minute,15*time.Minute)