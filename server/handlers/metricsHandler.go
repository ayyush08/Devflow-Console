package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/ayyush08/keploy-dashboard/config"
	"github.com/ayyush08/keploy-dashboard/models"
)

type PRResponse struct {
	State     string      `json:"state"`
	MergedAt  interface{} `json:"merged_at"`
	CreatedAt string      `json:"created_at"`
}

type RepoResponse struct {
	StargazersCount int    `json:"stargazers_count"`
	OpenIssuesCount int    `json:"open_issues_count"`
	ForksCount      int    `json:"forks_count"`
	PushedAt        string `json:"pushed_at"`
}



func FetchPRMetrics(owner, repo string) (models.PRMetrics, error) {

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/pulls?state=all?per_page=100"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+config.GetGithubToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.PRMetrics{}, fmt.Errorf("failed to fetch PR metrics: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.PRMetrics{}, fmt.Errorf("github API returned status %s", resp.Status)
	}

	fmt.Println("Response: ", resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return models.PRMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var prs []PRResponse
	if err := json.Unmarshal(body, &prs); err != nil {
		return models.PRMetrics{}, fmt.Errorf("failed to parse PRs: %v", err)
	}

	var totalMergeTime float64

	metrics := models.PRMetrics{
		Merged:   0,
		Pending:  0,
		Rejected: 0,
	}

	for _, pr := range prs {
		if pr.State == "closed" {
			if pr.MergedAt != nil {
				metrics.Merged++
				created, _ := time.Parse(time.RFC3339, pr.CreatedAt)
				merged, _ := time.Parse(time.RFC3339, pr.MergedAt.(string))
				totalMergeTime += merged.Sub(created).Hours() / 24
			} else {
				metrics.Rejected++
			}
		} else {
			metrics.Pending++
		}
	}

	metrics.PRCount = len(prs)

	if metrics.Merged > 0 {
		metrics.AvgMergeTime = math.Round((totalMergeTime / float64(metrics.Merged)) * 100) / 100
	}

	return metrics, nil
}

func FetchRepoMetrics(owner, repo string) (models.RepoMetrics, error) {
	url := "https://api.github.com/repos/" + owner + "/" + repo+"?per_page=100"

	log.Println("URL: ", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+config.GetGithubToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.RepoMetrics{}, fmt.Errorf("failed to fetch PR metrics: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return models.RepoMetrics{}, fmt.Errorf("GitHub API error: %s, response: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.RepoMetrics{}, fmt.Errorf("failed to read body: %v", err)
	}

	var repoData RepoResponse

	if err := json.Unmarshal(body, &repoData); err != nil {
		return models.RepoMetrics{}, fmt.Errorf("failed to parse repo: %v", err)
	}
	lastPushed, _ := time.Parse(time.RFC3339, repoData.PushedAt)
	daysSince := int(time.Since(lastPushed).Hours() / 24)

	log.Println("Repo Data: ", repoData)

	return models.RepoMetrics{
		Stars:       repoData.StargazersCount,
		OpenIssues:  repoData.OpenIssuesCount,
		Forks:       repoData.ForksCount,
		LastUpdated: fmt.Sprintf("%d days ago", daysSince),
	}, nil
}

func FetchTestMetrics(owner, repo string) (models.TestMetrics, error) {

	return models.TestMetrics{
		TestsPassed: 45,
		TestsFailed: 5,
		PassRate:    90.0,
	}, nil
}
