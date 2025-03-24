package models

type TileDataManager struct {
	TotalCommits 	   int `json:"totalCommits"`
	TotalPRs 		 int `json:"totalPRs"`
	TotalBugsOpen 	 int `json:"totalBugsOpen"`
	TotalTestsRun 	 int `json:"totalTestsRun"`
}

type AreaGraphDataManager struct {
	Date string `json:"date"`
	Commits int `json:"commits"`
	BugsReported int `json:"bugsReported"`
}

type BarGraphDataManager struct {
	Month string `json:"month"`
	PRsMerged int `json:"prsMerged"`
	BugsFixed int `json:"bugsFixed"`
}

type DonutChartDataManager struct {
	OpenPRs int `json:"openPRs"`
	MergedPRs int `json:"mergedPRs"`
	OpenBugs int `json:"openBugs"`
	ResolvedBugs int `json:"resolvedBugs"`
}


type ManagerMetrics struct {
	TileData TileDataManager `json:"tileData"`
	AreaGraphData []AreaGraphDataManager `json:"areaGraphData"`
	BarGraphData []BarGraphDataManager `json:"barGraphData"`
	DonutChartData DonutChartDataManager `json:"donutChartData"`
}




type ManagerMetricsGraphQLResponse struct {
	Repository struct {
		TotalCommits struct {
			Target struct {
				History struct {
					TotalCount int `json:"totalCount"`
				} `json:"history"`
			} `json:"target"`
		} `json:"totalCommits"`

		TotalPRs struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalPRs"`

		TotalBugsOpen struct {
			TotalCount int `json:"totalCount"`
		} `json:"totalBugsOpen"`

		TotalTestsRun struct {
			Target struct {
				History struct {
					TotalCount int `json:"totalCount"`
				} `json:"history"`
			} `json:"target"`
		} `json:"totalTestsRun"`

		
		CommitsHistory struct {
			Target struct {
				History struct {
					Edges []struct {
						Node struct {
							CommittedDate string `json:"committedDate"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"history"`
			} `json:"target"`
		} `json:"commitsHistory"`

		BugsReportedHistory struct {
			Nodes []struct {
				CreatedAt string `json:"createdAt"`
			} `json:"nodes"`
		} `json:"bugsReportedHistory"`

		
		PRsMergedHistory struct {
			Nodes []struct {
				MergedAt string `json:"mergedAt"`
			} `json:"nodes"`
		} `json:"prsMergedHistory"`

		BugsFixedHistory struct {
			Nodes []struct {
				ClosedAt string `json:"closedAt"`
			} `json:"nodes"`
		} `json:"bugsFixedHistory"`

		
		OpenPRs struct {
			TotalCount int `json:"totalCount"`
		} `json:"openPRs"`

		MergedPRs struct {
			TotalCount int `json:"totalCount"`
		} `json:"mergedPRs"`

		OpenBugs struct {
			TotalCount int `json:"totalCount"`
		} `json:"openBugs"`

		ResolvedBugs struct {
			TotalCount int `json:"totalCount"`
		} `json:"resolvedBugs"`
	} `json:"repository"`
}
