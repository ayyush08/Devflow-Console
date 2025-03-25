package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/ayyush08/devflow-console/config"
	"github.com/ayyush08/devflow-console/models"
	"github.com/ayyush08/devflow-console/queries"
)

func FetchDevMetrics(owner, repo string) (models.DeveloperMetrics, error) {
	cacheKey := fmt.Sprintf("metrics/developer/%s/%s", owner, repo)

	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		if item, ok := cachedItem.(config.MetricsCacheItem); ok && item.Type == "developer" {
			if metrics, ok := item.Value.(*models.DeveloperMetrics); ok {
				return *metrics, nil
			}
		}
	}

	sinceCommits := time.Now().AddDate(0, -6, 0).UTC().Format(time.RFC3339)

	graphQLPayload := models.GraphQLRequest{
		Query: queries.DeveloperMetricsQuery,
		Variables: map[string]string{
			"owner":        owner,
			"repo":         repo,
			"sinceCommits": sinceCommits,
		},
	}

	jsonPayload, err := json.Marshal(graphQLPayload)
	if err != nil {
		return models.DeveloperMetrics{}, fmt.Errorf("error marshalling graphql payload: %v", err)
	}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.DeveloperMetrics{}, fmt.Errorf("failed to fetch dev metrics: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return models.DeveloperMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var graphQLResponse models.DeveloperMetricsGraphQLResponse
	if err := json.Unmarshal(data, &graphQLResponse); err != nil {
		return models.DeveloperMetrics{}, fmt.Errorf("failed to parse GraphQL response: %v", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		return models.DeveloperMetrics{}, fmt.Errorf("error in fetching dev metrics: %v", graphQLResponse.Errors)
	}

	var devMetrics models.DeveloperMetrics
	repoData := graphQLResponse.Data.Repository

	devMetrics.TileData.TotalCommits = repoData.DefaultBranchRef.Target.History.TotalCount
	devMetrics.TileData.TotalPRs = repoData.PullRequests.TotalCount

	totalAdditions, totalDeletions, totalReviews := 0, 0, 0
	for _, edge := range repoData.Ref.Target.History.Edges {
		totalAdditions += edge.Node.Additions
		totalDeletions += edge.Node.Deletions
	}
	devMetrics.TileData.TotalLinesChanged = totalAdditions + totalDeletions

	for _, pr := range repoData.PullRequests.Nodes {
		totalReviews += pr.ReviewRequests.TotalCount
	}
	devMetrics.TileData.TotalReviewsReceived = totalReviews

	sincePRs := time.Now().AddDate(0, 0, -30).UTC().Format("2006-01-02")
	areaGraphDataMap := make(map[string]*models.AreaGraphDeveloperData)
	for _, edge := range repoData.Ref.Target.History.Edges {
		date := edge.Node.CommittedDate[:10]
		if _, exists := areaGraphDataMap[date]; !exists {
			areaGraphDataMap[date] = &models.AreaGraphDeveloperData{Date: date}
		}
		areaGraphDataMap[date].Commits++
	}
	for _, pr := range repoData.PullRequests.Nodes {
		date := pr.CreatedAt[:10]
		if _, exists := areaGraphDataMap[date]; !exists {
			areaGraphDataMap[date] = &models.AreaGraphDeveloperData{Date: date}
		}
		areaGraphDataMap[date].PullRequests++
	}

	areaGraphData := make([]models.AreaGraphDeveloperData, 0, len(areaGraphDataMap))
	for _, data := range areaGraphDataMap {
		if data.Date >= sincePRs {
			areaGraphData = append(areaGraphData, *data)
		}
	}
	sort.Slice(areaGraphData, func(i, j int) bool {
		return areaGraphData[i].Date > areaGraphData[j].Date
	})
	devMetrics.AreaGraphData = areaGraphData

	barGraphDataMap := make(map[string]*models.BarGraphDeveloperData)
	for _, edge := range repoData.Ref.Target.History.Edges {
		yearMonth := edge.Node.CommittedDate[:7]
		t, _ := time.Parse("2006-01", yearMonth)
		monthName := t.Format("January")

		if _, exists := barGraphDataMap[monthName]; !exists {
			barGraphDataMap[monthName] = &models.BarGraphDeveloperData{Month: monthName}
		}
		barGraphDataMap[monthName].Additions += edge.Node.Additions
		barGraphDataMap[monthName].Deletions += edge.Node.Deletions
	}

	barGraphData := make([]models.BarGraphDeveloperData, 0, len(barGraphDataMap))
	for _, data := range barGraphDataMap {
		barGraphData = append(barGraphData, *data)
	}
	sort.Slice(barGraphData, func(i, j int) bool {
		t1, _ := time.Parse("January", barGraphData[i].Month)
		t2, _ := time.Parse("January", barGraphData[j].Month)
		return t1.Month() > t2.Month()
	})
	devMetrics.BarGraphData = barGraphData

	openPRs, mergedPRs, closedPRs, pendingReviews := 0, 0, 0, 0
	for _, pr := range repoData.PullRequests.Nodes {
		switch pr.State {
		case "OPEN":
			openPRs++
		case "MERGED":
			mergedPRs++
		case "CLOSED":
			closedPRs++
		}
		pendingReviews += pr.ReviewRequests.TotalCount
	}
	devMetrics.DonutChartData = models.DonutChartDeveloperData{
		MergedPRs:      mergedPRs,
		ClosedPRs:      closedPRs,
		OpenPRs:        openPRs,
		PendingReviews: pendingReviews,
	}

	cacheItem := config.MetricsCacheItem{
		Type:  "developer",
		Value: &devMetrics,
	}
	config.GlobalCache.Set(cacheKey, cacheItem, 0)

	return devMetrics, nil
}
