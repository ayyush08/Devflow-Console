package utils

import (
	"sort"
	"time"

	"github.com/ayyush08/devflow-console/models"
)

func GenerateAreaGraphData(prTimestamps []models.PRNode, commitTimestamps []models.CommitNode) []models.AreaGraphGeneralData {
	prs := make(map[string]models.AreaGraphGeneralData)

	for _, t := range prTimestamps {
		date := t.CreatedAt[:10]
		point := prs[date]
		point.Date = date
		point.PullRequests++
		prs[date] = point
	}

	for _, t := range commitTimestamps {

		date := t.CommittedDate[:10]
		point := prs[date]
		point.Date = date
		point.Commits++
		prs[date] = point
	}

	var result []models.AreaGraphGeneralData
	for _, v := range prs {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		t1, _ := time.Parse("2006-01-02", result[i].Date)
		t2, _ := time.Parse("2006-01-02", result[j].Date)
		return t1.After(t2) // Sort in descending order
	})

	return result
}

func GenerateBarGraphData(commits []struct {
	Node struct {
		CommittedDate string `json:"committedDate"`
		Additions     int    `json:"additions"`
		Deletions     int    `json:"deletions"`
	} `json:"node"`
}) []models.BarGraphGeneralData {
	monthlyData := make(map[string]models.BarGraphGeneralData)

	for _, commit := range commits {
		date, err := time.Parse(time.RFC3339, commit.Node.CommittedDate)
		if err != nil {
			continue
		}
		month := date.Format("January")

		entry := monthlyData[month]
		entry.Month = month
		entry.Additions += commit.Node.Additions
		entry.Deletions += commit.Node.Deletions
		monthlyData[month] = entry
	}

	var barGraphData []models.BarGraphGeneralData
	for _, data := range monthlyData {
		barGraphData = append(barGraphData, data)
	}

	monthOrder := map[string]int{
		"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6,
		"July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12,
	}

	sort.Slice(barGraphData, func(i, j int) bool {
		return monthOrder[barGraphData[i].Month] < monthOrder[barGraphData[j].Month]
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
