package config

import "os"

func GetGithubToken() string {
	return os.Getenv("GITHUB_ACCESS_TOKEN")
}
