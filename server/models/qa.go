package models

type TileDataQa struct{
	TotalTestsRun int `json:"totalTestsRun"`
	TotalBugsReported int `json:"totalBugsReported"`
	TotalBugsResolved int `json:"totalBugsResolved"`
	TotalTestSuites int `json:"totalTestSuites"`
}

type AreaGraphQaData struct{
	Date string `json:"date"`
	TestsRun int `json:"testsRun"`
	BugsReported int `json:"bugsReported"`
}

type BarGraphQaData struct{
	Month string `json:"month"`
	BugsFixed int `json:"bugsFixed"`
	TestExecutions int `json:"testExecutions"`
}

type DonutChartQaData struct{
	SuccessTests int `json:"successfullTests"`
	FailedTests int `json:"failedTests"`
	SkippedTest int `json:"skippedTests"`
}

type QaMetrics struct{
	TileData TileDataQa `json:"tileData"`
	AreaGraphData []AreaGraphQaData `json:"areaGraphData"`
	BarGraphData []BarGraphQaData `json:"barGraphData"`
	DonutChartData DonutChartQaData `json:"donutChartData"`
}


type QaMetricsGraphQLResponse struct {
	Repository struct {
		TotalBugsReported struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalBugsReported"`

		TotalBugsResolved struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalBugsResolved"`

		TotalTestSuites struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalTestSuites"`

		
		Issues struct {
			Nodes []struct {
				CreatedAt string `json:"createdAt"`
			} `json:"nodes"`
		} `json:"issues"`

		
		ClosedIssues struct {
			Nodes []struct {
				ClosedAt string `json:"closedAt"`
			} `json:"nodes"`
		} `json:"closedIssues"`

		
		DefaultBranchRef struct {
			Target struct {
				CheckSuites struct {
					Nodes []struct {
						Conclusion string `json:"conclusion"`
					} `json:"nodes"`
				} `json:"checkSuites"`
			} `json:"target"`
		} `json:"defaultBranchRef"`
	} `json:"repository"`
}
