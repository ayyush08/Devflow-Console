package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ayyush08/keploy-dashboard/config"
	"github.com/ayyush08/keploy-dashboard/models"
)

type PRResponse struct {
	State    string      `json:"state"`
	MergedAt interface{} `json:"merged_at"` 
}

func FetchPRMetrics(owner, repo string) (models.PRMetrics, error) {
	token := config.GetGithubToken()

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/pulls?state=all"

	log.Println("URL: ", url)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.PRMetrics{}, fmt.Errorf("failed to fetch PR metrics: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return models.PRMetrics{}, fmt.Errorf("github API returned status %s", resp.Status)
	}

	fmt.Println("Response: ",resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return models.PRMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var prs []PRResponse
	if err := json.Unmarshal(body, &prs); err != nil {
		return models.PRMetrics{}, fmt.Errorf("failed to parse PRs: %v", err)
	}

	metrics := models.PRMetrics{
		Merged:  0,
		Pending: 0,
		Rejected: 0,
	}

	for _, pr := range prs {
		switch pr.State{
		case "closed":
			if pr.MergedAt != nil {
				metrics.Merged++
			} else {
				metrics.Rejected++
			}
		case "open":
			metrics.Pending++
		}
	}
	return metrics, nil
}


