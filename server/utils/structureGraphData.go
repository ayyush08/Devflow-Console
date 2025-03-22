package utils

import (
	"sort"
	"time"

	"github.com/ayyush08/devflow-console/models"
)

func GenerateAreaGraphData(prTimestamps []models.PRNode, commitTimestamps []models.CommitNode) []models.AreaGraphData {
	prs := make(map[string]models.AreaGraphData)

	// Process PR timestamps
	for _, t := range prTimestamps {
		date := t.CreatedAt[:10] // Extract "YYYY-MM-DD"
		point := prs[date]
		point.Date = date
		point.PullRequests++
		prs[date] = point
	}

	// Process Commit timestamps
	for _, t := range commitTimestamps {

		date := t.CommittedDate[:10]
		point := prs[date]
		point.Date = date
		point.Commits++
		prs[date] = point
	}

	// Convert map to slice
	var result []models.AreaGraphData
	for _, v := range prs {
		result = append(result, v)
	}

	return result
}

func GenerateBarGraphData(commits []struct {
	Node struct {
		CommittedDate string `json:"committedDate"`
		Additions     int    `json:"additions"`
		Deletions     int    `json:"deletions"`
	} `json:"node"`
}) []models.BarGraphData {
	monthlyData := make(map[string]models.BarGraphData)

	for _, commit := range commits {
		date, err := time.Parse(time.RFC3339, commit.Node.CommittedDate)
		if err != nil {
			continue
		}
		month := date.Format("2006-01") // "YYYY-MM"

		// Aggregate additions and deletions per month
		entry := monthlyData[month]
		entry.Month = month
		entry.Additions += commit.Node.Additions
		entry.Deletions += commit.Node.Deletions
		monthlyData[month] = entry
	}

	// Convert map to slice
	var barGraphData []models.BarGraphData
	for _, data := range monthlyData {
		barGraphData = append(barGraphData, data)
	}

	// Sort by month
	sort.Slice(barGraphData, func(i, j int) bool {
		return barGraphData[i].Month < barGraphData[j].Month
	})

	return barGraphData
}

func ExtractPRTimestamps(edges []struct{ Node struct{ CreatedAt string } }) []string {
	timestamps := make([]string, len(edges))
	for i, edge := range edges {
		timestamps[i] = edge.Node.CreatedAt
	}
	return timestamps
}

func ExtractCommitTimestamps(edges []struct {
	Node struct{ CommittedDate string }
}) []string {
	timestamps := make([]string, len(edges))
	for i, edge := range edges {
		timestamps[i] = edge.Node.CommittedDate
	}
	return timestamps
}
