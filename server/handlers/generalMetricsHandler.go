package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/ayyush08/devflow-console/server/config"
	"github.com/ayyush08/devflow-console/server/models"
	"github.com/ayyush08/devflow-console/server/queries"
)

func FetchGeneralMetrics(owner string, repo string) (models.GeneralMetrics, error) {
	cacheKey := fmt.Sprintf("general/%s/%s", owner, repo)

	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		item, ok := cachedItem.(config.MetricsCacheItem)
		if !ok || item.Type != "general" {
			return models.GeneralMetrics{}, fmt.Errorf("cache type mismatch: expected general, got %T", cachedItem)
		}

		generalMetrics, ok := item.Value.(*models.GeneralMetrics)
		if !ok {
			return models.GeneralMetrics{}, fmt.Errorf("cache value type mismatch")
		}

		return *generalMetrics, nil
	}

	since := time.Now().AddDate(0, -6, 0).UTC().Format(time.RFC3339)

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

	repoData := graphQLResponse.Data.Repository
	var generalMetrics models.GeneralMetrics

	generalMetrics.TileData.TotalCommits = repoData.TotalCommits.Target.History.TotalCount
	generalMetrics.TileData.TotalIssues = repoData.Issues.TotalCount
	generalMetrics.TileData.TotalPRs = repoData.PullRequests.TotalCount
	generalMetrics.TileData.TotalStars = repoData.StargazerCount

	areaGraphDataMap := make(map[string]*models.AreaGraphGeneralData)

	cutoffTime := time.Now().AddDate(0,0,-30).Format("2006-01-02")

	for _, pr := range repoData.PullRequests.Edges {
		date := pr.Node.CreatedAt[:10]
		
		if date >= cutoffTime {
			if _, ok := areaGraphDataMap[date]; !ok{
				areaGraphDataMap[date] = &models.AreaGraphGeneralData{
					Date: date,
				}
			}
			areaGraphDataMap[date].PullRequests++
		}

	}
	for _, commit := range repoData.RecentCommits.Target.History.Edges {
		date := commit.Node.CommittedDate[:10]
		
		if date >= cutoffTime{
			if _,ok := areaGraphDataMap[date]; !ok{
				areaGraphDataMap[date] = &models.AreaGraphGeneralData{
					Date: date,
				}
			}
			areaGraphDataMap[date].Commits++
		}

	}
	areaGraphData := make([]models.AreaGraphGeneralData, 0, len(areaGraphDataMap))
	for _, data := range areaGraphDataMap {
		areaGraphData = append(areaGraphData, *data)
	}

	sort.Slice(areaGraphData, func(i, j int) bool {
		return areaGraphData[i].Date > areaGraphData[j].Date
	})
	generalMetrics.AreaGraphData = areaGraphData

	barGraphDataMap := make(map[string]models.BarGraphGeneralData)
	for _, commit := range repoData.BarData.Target.History.Edges {
		date, err := time.Parse(time.RFC3339, commit.Node.CommittedDate)
		if err != nil {
			continue
		}
		month := date.Format("January")
		entry := barGraphDataMap[month]
		entry.Month = month
		entry.Additions += commit.Node.Additions
		entry.Deletions += commit.Node.Deletions
		barGraphDataMap[month] = entry
	}
	var barGraphData []models.BarGraphGeneralData
	for _, data := range barGraphDataMap {
		barGraphData = append(barGraphData, data)
	}
	sort.Slice(barGraphData, func(i, j int) bool {
		monthOrder := map[string]int{
			"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6,
			"July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12,
		}
		return monthOrder[barGraphData[i].Month] < monthOrder[barGraphData[j].Month]
	})
	generalMetrics.BarGraphData = barGraphData


	generalMetrics.DonutChartData.ClosedPRs = repoData.ClosedPRs.TotalCount
	generalMetrics.DonutChartData.MergedPRs = repoData.MergedPRs.TotalCount
	generalMetrics.DonutChartData.OpenPRs = repoData.OpenPRs.TotalCount

	cacheItem := config.MetricsCacheItem{
		Type:  "general",
		Value: &generalMetrics,
	}
	config.GlobalCache.Set(cacheKey, cacheItem, 0)

	return generalMetrics, nil
}