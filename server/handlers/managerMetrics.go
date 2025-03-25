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

func FetchManagerTemplate(owner, repo string) (models.ManagerMetrics, error) {
	cacheKey := fmt.Sprintf("metrics/manager/%s/%s", owner, repo)

	if cachedItem, found := config.GlobalCache.Get(cacheKey); found {
		if item, ok := cachedItem.(config.MetricsCacheItem); ok && item.Type == "manager" {
			if metrics, ok := item.Value.(*models.ManagerMetrics); ok {
				return *metrics, nil
			}
		}
	}

	graphQLPayload := models.GraphQLRequest{
		Query: queries.ManagerMetricsQuery,
		Variables: map[string]string{
			"owner": owner,
			"repo":  repo,
		},
	}

	jsonPayload, err := json.Marshal(graphQLPayload)

	if err != nil {
		return models.ManagerMetrics{}, fmt.Errorf("error marshalling graphql payload: %v", err)
	}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.ManagerMetrics{}, fmt.Errorf("failed to fetch qa metrics: %v", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return models.ManagerMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var graphQLResponse models.ManagerMetricsGraphQLResponse
	if err := json.Unmarshal(data, &graphQLResponse); err != nil {
		return models.ManagerMetrics{}, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		return models.ManagerMetrics{}, fmt.Errorf("failed to fetch qa metrics: %v", graphQLResponse.Errors)
	}

	var managerMetrics models.ManagerMetrics

	repoData := graphQLResponse.Data.Repository

	managerMetrics.TileData.TotalBugsOpen = repoData.TotalBugsOpen.TotalCount
	managerMetrics.TileData.TotalCommits = repoData.TotalCommits.Target.History.TotalCount
	managerMetrics.TileData.TotalIssuesOpen = repoData.TotalIssuesOpen.TotalCount
	managerMetrics.TileData.TotalPRs = repoData.TotalPRs.TotalCount

	areaGraphDataMap := make(map[string]*models.AreaGraphDataManager)

	cutoffTime := time.Now().AddDate(0, 0, -30).Format("2006-01-02")

	for _, commit := range repoData.CommitsHistory.Target.History.Edges {
		date := commit.Node.CommittedDate[:10]

		if date >= cutoffTime {
			if _, ok := areaGraphDataMap[date]; !ok {
				areaGraphDataMap[date] = &models.AreaGraphDataManager{
					Date: date,
				}
			}
			areaGraphDataMap[date].Commits++
		}
	}

	for _, bug := range repoData.BugsReportedHistory.Nodes {
		date := bug.CreatedAt[:10]

		if date >= cutoffTime {
			if _, ok := areaGraphDataMap[date]; !ok {
				areaGraphDataMap[date] = &models.AreaGraphDataManager{
					Date: date,
				}
			}

			areaGraphDataMap[date].BugsReported++
		}
	}

	areaGraphData := make([]models.AreaGraphDataManager, 0, len(areaGraphDataMap))
	for _, data := range areaGraphDataMap {
		areaGraphData = append(areaGraphData, *data)
	}

	sort.Slice(areaGraphData, func(i, j int) bool {
		return areaGraphData[i].Date > areaGraphData[j].Date
	})

	barGraphDataMap := make(map[string]*models.BarGraphDataManager)

	cutoffMonth := time.Now().AddDate(0, -6, 0).Format("2006-01")
	for _, pr := range repoData.PRsMergedHistory.Nodes {
		month := pr.MergedAt[:7]
		t, _ := time.Parse("2006-01", month)
		monthName := t.Format("January")
		if month >= cutoffMonth {
			if _, exists := barGraphDataMap[monthName]; !exists {
				barGraphDataMap[monthName] = &models.BarGraphDataManager{
					Month: monthName,
				}
			}
			barGraphDataMap[monthName].PRsMerged++
		}
	}

	for _, bug := range repoData.BugsFixedHistory.Nodes {
		month := bug.ClosedAt[:7]
		t, _ := time.Parse("2006-01", month)
		monthName := t.Format("January")

		if month >= cutoffMonth {
			if _, exists := barGraphDataMap[monthName]; !exists {
				barGraphDataMap[monthName] = &models.BarGraphDataManager{
					Month: monthName,
				}
			}
			barGraphDataMap[monthName].BugsFixed++
		}

	}

	// log.Println(barGraphDataMap)
	barGraphData := make([]models.BarGraphDataManager, 0, len(barGraphDataMap))

	for _, data := range barGraphDataMap {
		barGraphData = append(barGraphData, *data)
	}
	sort.Slice(barGraphData, func(i, j int) bool {
		return barGraphData[i].Month > barGraphData[j].Month
	})

	managerMetrics.BarGraphData = barGraphData
	managerMetrics.AreaGraphData = areaGraphData

	managerMetrics.DonutChartData.MergedPRs = repoData.MergedPRs.TotalCount
	managerMetrics.DonutChartData.OpenBugs = repoData.OpenBugs.TotalCount
	managerMetrics.DonutChartData.OpenPRs = repoData.OpenPRs.TotalCount
	managerMetrics.DonutChartData.ResolvedBugs = repoData.ResolvedBugs.TotalCount

	cacheItem := config.MetricsCacheItem{
		Type:  "manager",
		Value: &managerMetrics,
	}

	config.GlobalCache.Set(cacheKey, cacheItem, 0)

	return managerMetrics, nil

}
