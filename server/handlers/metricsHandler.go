package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ayyush08/devflow-console/config"
	"github.com/ayyush08/devflow-console/models"
	"github.com/ayyush08/devflow-console/queries"
	"github.com/ayyush08/devflow-console/utils"
)

func FetchMetrics(owner string, repo string, template string) (models.DashboardMetrics, error) {

	cacheKey := fmt.Sprintf("metrics/%s/%s/%s", template, owner, repo)

	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		item, ok := cachedItem.(config.MetricsCacheItem)
		if !ok || item.Type != "template" {
			return models.DashboardMetrics{}, fmt.Errorf("cache type mismatch: expected dashboard, got %T", cachedItem)
		}
	
		metrics, ok := item.Value.(*models.DashboardMetrics) // Use pointer type assertion
		if !ok {
			return models.DashboardMetrics{}, fmt.Errorf("cache value type mismatch")
		}
	
		return *metrics, nil // Dereference before returning
	}
	

	graphQLPayload := models.GraphQLRequest{
		Query: queries.MetricsQuery,
		Variables: map[string]string{
			"owner": owner,
			"name":  repo,
		},
	}

	jsonPayload, err := json.Marshal(graphQLPayload)

	if err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("error marshalling graphql payload: %v", err)
	}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("failed to fetch metrics: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var graphQLResponse models.GraphQLResponse

	if err := json.Unmarshal(data, &graphQLResponse); err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("failed to parse GraphQL response: %v", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		return models.DashboardMetrics{}, fmt.Errorf("GraphQL error: %v", graphQLResponse.Errors[0].Message)
	}

	var metrics models.DashboardMetrics

	metrics.PRMetrics = utils.ExtractPRMetrics(graphQLResponse)
	metrics.RepoMetrics = utils.ExtractRepoMetrics(graphQLResponse)
	metrics.TestMetrics = utils.ExtractTestMetrics(graphQLResponse)

	cacheItem := config.MetricsCacheItem{
		Type:  "template",
		Value: &metrics,
	}

	config.GlobalCache.Set(cacheKey, cacheItem, 0)

	return metrics, nil

}

func FetchGeneralMetrics(owner string, repo string) (models.GeneralMetrics, error) {
	cacheKey := fmt.Sprintf("general/%s/%s", owner, repo)

	

	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		item, ok := cachedItem.(config.MetricsCacheItem)
		if !ok || item.Type != "general" {
			return models.GeneralMetrics{}, fmt.Errorf("cache type mismatch: expected general, got %T", cachedItem)
		}
	
		generalMetrics, ok := item.Value.(*models.GeneralMetrics) // Use pointer type assertion
		if !ok {
			return models.GeneralMetrics{}, fmt.Errorf("cache value type mismatch")
		}
	
		return *generalMetrics, nil // Dereference before returning
	}
	

	since := time.Now().AddDate(0, 0, -30).UTC().Format(time.RFC3339)

	graphQLPayload := models.GraphQLRequest{
		Query: queries.GeneralMetricsQuery,
		Variables: map[string]string{
			"owner": owner,
			"repo":  repo,
			"since": since,
		},
	}

	jsonPayload, err := json.Marshal(graphQLPayload)

	if err != nil {
		return models.GeneralMetrics{}, fmt.Errorf("error marshalling graphql payload: %v", err)
	}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return models.GeneralMetrics{}, fmt.Errorf("failed to fetch general metrics: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return models.GeneralMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var graphQLResponse models.GeneralMetricsGraphQLResponse

	if err := json.Unmarshal(data, &graphQLResponse); err != nil {
		return models.GeneralMetrics{}, fmt.Errorf("failed to parse GraphQL response: %v", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		return models.GeneralMetrics{}, fmt.Errorf("GraphQL error: %v", graphQLResponse.Errors[0].Message)
	}

	var generalMetrics models.GeneralMetrics

	repoData := graphQLResponse.Data.Repository

	generalMetrics.TotalCommits = repoData.TotalCommits.Target.History.TotalCount
	generalMetrics.TotalIssues = repoData.Issues.TotalCount
	generalMetrics.TotalPRs = repoData.PullRequests.TotalCount
	generalMetrics.TotalStars = repoData.StargazerCount

	prs := repoData.PullRequests.Edges
	commits := repoData.RecentCommits.Target.History.Edges

	var prTimestamps []models.PRNode
	for _, edge := range prs {
		prTimestamps = append(prTimestamps, models.PRNode{CreatedAt: edge.Node.CreatedAt})
	}

	var commitTimestamps []models.CommitNode
	for _, edge := range commits {
		commitTimestamps = append(commitTimestamps, models.CommitNode{CommittedDate: edge.Node.CommittedDate})
	}

	generalMetrics.AreaGraphData = utils.GenerateAreaGraphData(prTimestamps, commitTimestamps)

	generalMetrics.DonutChartData.ClosedPRs = repoData.ClosedPRs.TotalCount
	generalMetrics.DonutChartData.MergedPRs = repoData.MergedPRs.TotalCount
	generalMetrics.DonutChartData.OpenPRs = repoData.OpenPRs.TotalCount

	barData := repoData.BarData.Target.History.Edges

	generalMetrics.BarGraphData = utils.GenerateBarGraphData(barData)


	cacheItem := config.MetricsCacheItem{
		Type:  "general",
		Value: &generalMetrics,
	}

	config.GlobalCache.Set(cacheKey, cacheItem, 0)

	return generalMetrics, nil
}
