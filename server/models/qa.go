package models

type TileDataQa struct{
	TotalBugsReported int `json:"totalBugsReported"`
	TotalBugsResolved int `json:"totalBugsResolved"`
	TotalDiscussions int `json:"totalDiscussions"`
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
	Data struct {
	Repository struct {
		TotalBugsReported struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalBugsReported"`

		TotalBugsResolved struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalBugsResolved"`

		TotalDiscussions struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalDiscussions"`

		
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
						CreatedAt string `json:"createdAt"`
						CheckRuns struct {
							TotalCount int `json:"totalCount"`
						} `json:"checkRuns"`
					} `json:"nodes"`
					
				} `json:"checkSuites"`
			} `json:"target"`
		} `json:"defaultBranchRef"`
	} `json:"repository"`
					} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}
