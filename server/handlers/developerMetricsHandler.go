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
)

func FetchDevMetrics(owner string, repo string) (models.DeveloperMetrics, error) {
	cacheKey := fmt.Sprintf("metrics/developer/%s/%s", owner, repo)

	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		item, ok := cachedItem.(config.MetricsCacheItem)
		if !ok || item.Type != "developer" {
			return models.DeveloperMetrics{}, fmt.Errorf("cache type mismatch: expected dashboard, got %T", cachedItem)
		}

		metrics, ok := item.Value.(*models.DeveloperMetrics)
		if !ok {
			return models.DeveloperMetrics{}, fmt.Errorf("cache value type mismatch")
		}

		return *metrics, nil
	}
	since := time.Now().AddDate(0, 0, -30).UTC().Format(time.RFC3339)
	graphQLPayload := models.GraphQLRequest{
		Query: queries.DeveloperMetricsQuery,
		Variables: map[string]string{
			"owner": owner,
			"name":  repo,
			"since": since,
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

	if(len(graphQLResponse.Errors) > 0){
		return models.DeveloperMetrics{}, fmt.Errorf("error in fetching dev metrics: %v", graphQLResponse.Errors)
	}

	var devMetrics models.DeveloperMetrics

	repoData := graphQLResponse.Repository

	devMetrics.TileData.TotalCommits = repoData.DefaultBranchRef.Target.History.TotalCount
	
	totalAdditions, totalDeletions := 0, 0

	for _,edge := range repoData.Ref.Target.History.Edges{
		totalAdditions+= edge.Node.Additions
		totalDeletions+= edge.Node.Deletions
	}
	devMetrics.TileData.TotalLinesChanged = totalAdditions+totalDeletions

	devMetrics.TileData.TotalPRs = repoData.PullRequests.TotalCount
	totalReviews := 0
	for _, pr := range repoData.PullRequests.Nodes {
		totalReviews += pr.ReviewRequests.TotalCount
	}
	devMetrics.TileData.TotalReviewsReceived = totalReviews

	areaGraphDataMap := make(map[string]*models.AreaGraphDeveloperData)
	for _, edge := range repoData.Ref.Target.History.Edges {
		date := edge.Node.CommittedDate[:10] // Extract YYYY-MM-DD
		if _, exists := areaGraphDataMap[date]; !exists {
			areaGraphDataMap[date] = &models.AreaGraphDeveloperData{Date: date}
		}
		areaGraphDataMap[date].Commits++
	}
	for _, pr := range repoData.PullRequests.Nodes {
		date := pr.CreatedAt[:10] // Extract YYYY-MM-DD
		if _, exists := areaGraphDataMap[date]; !exists {
			areaGraphDataMap[date] = &models.AreaGraphDeveloperData{Date: date}
		}
		areaGraphDataMap[date].PullRequests++
	}
	for _, data := range areaGraphDataMap {
		devMetrics.AreaGraphData = append(devMetrics.AreaGraphData, *data)
	}

	barGraphDataMap := make(map[string]*models.BarGraphDeveloperData)
	for _, edge := range repoData.Ref.Target.History.Edges {
		month := edge.Node.CommittedDate[:7] // Extract YYYY-MM
		if _, exists := barGraphDataMap[month]; !exists {
			barGraphDataMap[month] = &models.BarGraphDeveloperData{Month: month}
		}
		barGraphDataMap[month].Additions += edge.Node.Additions
		barGraphDataMap[month].Deletions += edge.Node.Deletions
	}
	for _, data := range barGraphDataMap {
		devMetrics.BarGraphData = append(devMetrics.BarGraphData, *data)
	}

	openPRs, mergedPRs, closedPRs := 0, 0, 0
	pendingReviews := 0
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
		Type: "developer",
		Value: &devMetrics,
	}

	config.GlobalCache.Set(cacheKey, cacheItem,0)

	

	




	return devMetrics, nil
}
