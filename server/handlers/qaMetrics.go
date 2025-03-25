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

func FetchQaMetrics(owner, repo string) (models.QaMetrics, error) {

	cacheKey := fmt.Sprintf("metrics/qa/%s/%s", owner, repo)
	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		if item, ok := cachedItem.(config.MetricsCacheItem); ok && item.Type == "developer" {
			if metrics, ok := item.Value.(*models.QaMetrics); ok {
				return *metrics, nil
			}
		}
	}

	graphQLPayload := models.GraphQLRequest{
		Query: queries.QaMetricsQuery,
		Variables: map[string]string{
			"owner": owner,
			"repo":  repo,
		},
	}

	jsonPayload, err := json.Marshal(graphQLPayload)
	if err != nil {
		return models.QaMetrics{}, fmt.Errorf("error marshalling graphql payload: %v", err)
	}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.QaMetrics{}, fmt.Errorf("failed to fetch qa metrics: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return models.QaMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var graphQLResponse models.QaMetricsGraphQLResponse
	if err := json.Unmarshal(data, &graphQLResponse); err != nil {
		return models.QaMetrics{}, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		return models.QaMetrics{}, fmt.Errorf("failed to fetch qa metrics: %v", graphQLResponse.Errors)
	}

	var qaMetrics models.QaMetrics

	repoData := graphQLResponse.Data.Repository

	qaMetrics.TileData.TotalBugsReported = repoData.TotalBugsReported.TotalCount
	qaMetrics.TileData.TotalBugsResolved = repoData.TotalBugsResolved.TotalCount
	qaMetrics.TileData.TotalDiscussions = repoData.TotalDiscussions.TotalCount

	areaGraphDataMap := make(map[string]*models.AreaGraphQaData)

	cutoffTime := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	for _, issue := range repoData.Issues.Nodes {
		date := issue.CreatedAt[:10]

		if date >= cutoffTime {
			if _, exists := areaGraphDataMap[date]; !exists {
				areaGraphDataMap[date] = &models.AreaGraphQaData{Date: date}
			}
			areaGraphDataMap[date].BugsReported++
		}
	}

	for _, suite := range repoData.DefaultBranchRef.Target.CheckSuites.Nodes {
		date := suite.CreatedAt[:10]
		if date >= cutoffTime {
			if _, exists := areaGraphDataMap[date]; !exists {
				areaGraphDataMap[date] = &models.AreaGraphQaData{Date: date}
			}
			areaGraphDataMap[date].TestsRun++
		}
	}

	areaGraphData := make([]models.AreaGraphQaData, 0, len(areaGraphDataMap))
	for _, data := range areaGraphDataMap {
		areaGraphData = append(areaGraphData, *data)
	}
	sort.Slice(areaGraphData, func(i, j int) bool {
		return areaGraphData[i].Date > areaGraphData[j].Date
	})

	barGraphDataMap := make(map[string]*models.BarGraphQaData)
	cutoffMonth := time.Now().AddDate(0, -6, 0).Format("2006-01")

	for _, issue := range repoData.ClosedIssues.Nodes {
		month := issue.ClosedAt[:7] 
		t, _ := time.Parse("2006-01", month)
		monthName := t.Format("January")
		if month >= cutoffMonth {
			if _, exists := barGraphDataMap[monthName]; !exists {
				barGraphDataMap[monthName] = &models.BarGraphQaData{Month: monthName}
			}
			barGraphDataMap[monthName].BugsFixed++
		}
	}

	for _, test := range repoData.DefaultBranchRef.Target.CheckSuites.Nodes{
		month := test.CreatedAt[:7]
		t, _ := time.Parse("2006-01", month)
		monthName := t.Format("January")

		if month >= cutoffMonth {
			if _, exists := barGraphDataMap[monthName]; !exists{
				barGraphDataMap[monthName] = &models.BarGraphQaData{
					Month: monthName,
				}
			}
			barGraphDataMap[monthName].TestExecutions++
		}
		
	}

	barGraphData := make([]models.BarGraphQaData, 0, len(barGraphDataMap))
	for _, data := range barGraphDataMap {
		barGraphData = append(barGraphData, *data)
	}
	sort.Slice(barGraphData, func(i, j int) bool {
		return barGraphData[i].Month > barGraphData[j].Month
	})
	
	donutChartData := models.DonutChartQaData{}

	for _, suite := range repoData.DefaultBranchRef.Target.CheckSuites.Nodes {
		switch suite.Conclusion {
		case "SUCCESS":
			donutChartData.SuccessTests++
		case "FAILURE":
			donutChartData.FailedTests++
		case "SKIPPED":
			donutChartData.SkippedTest++
		}
	}

	qaMetrics.AreaGraphData = areaGraphData
	qaMetrics.BarGraphData = barGraphData
	qaMetrics.DonutChartData = donutChartData

	cacheItem := config.MetricsCacheItem{
		Type:  "qa",
		Value: &qaMetrics,
	}

	config.GlobalCache.Set(cacheKey, cacheItem, 0)

	return qaMetrics, nil
}
