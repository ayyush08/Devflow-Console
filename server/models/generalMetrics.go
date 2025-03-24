package models


type AreaGraphGeneralData struct {
	Date string `json:"date"`
	Commits int `json:"commits"`
	PullRequests int `json:"pullRequests"`
}

type BarGraphGeneralData struct {
	Month string `json:"month"`
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
}

type DonutChartGeneralData struct {
	MergedPRs int `json:"merged"`
	ClosedPRs int `json:"closed"`
	OpenPRs int `json:"open"`
}

type GeneralMetrics struct {
	TileData struct {
		TotalPRs int `json:"totalPRs"`
		TotalCommits int `json:"totalCommits"`
		TotalIssues int `json:"totalIssues"`
		TotalStars int `json:"totalStars"`
	} `json:"tileData"`
	AreaGraphData []AreaGraphGeneralData `json:"areaGraphData"`
	BarGraphData []BarGraphGeneralData `json:"barGraphData"`
	DonutChartData DonutChartGeneralData `json:"donutChartData"`
}