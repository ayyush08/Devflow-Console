package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ayyush08/keploy-dashboard/config"
	"github.com/ayyush08/keploy-dashboard/models"
)

func FetchPRMetrics(owner, repo string) (models.PRMetrics, error) {
	token := config.GetGithubToken()

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/pulls?state=all"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.PRMetrics{}, err
	}
	defer resp.Body.Close()

	fmt

	body, _ := io.ReadAll(resp.Body)
	var prs []map[string]interface{}
	json.Unmarshal(body, &prs)

	merged, pending, rejected := 0, 0, 0
	for _, pr := range prs {
		if pr["merged_at"] != nil {
			merged++
		} else if pr["state"] == "open" {
			pending++
		} else {
			rejected++
		}
	}
	return models.PRMetrics{Merged: merged, Pending: pending, Rejected: rejected}, nil

}
