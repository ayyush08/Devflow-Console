package models

type TileDataDeveloper struct {
	TotalPRs             int `json:"totalPRs"`
	TotalCommits         int `json:"totalCommits"`
	TotalLinesChanged    int `json:"totalLinesChanged"`
	TotalReviewsReceived int `json:"totalReviewsReceived"`
}

type AreaGraphDeveloperData struct {
	Date string `json:"date"`
	Commits int `json:"dailyCommits"`
	PullRequests int `json:"dailyPullRequests"`
}

type BarGraphDeveloperData struct {
	Month string `json:"month"`
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
}
type DonutChartDeveloperData struct {
	MergedPRs int `json:"mergedPRs"`
	ClosedPRs int `json:"closedPRs"`
	OpenPRs int `json:"openPRs"`
	PendingReviews int `json:"pendingReviews"`
}

type DeveloperMetrics struct {
	TileData TileDataDeveloper `json:"tileData"`
	AreaGraphData []AreaGraphDeveloperData `json:"areaGraphData"`
	BarGraphData []BarGraphDeveloperData `json:"barGraphData"`
	DonutChartData DonutChartDeveloperData `json:"donutChartData"`
}

type DeveloperMetricsGraphQLResponse struct {
	Repository struct {
		DefaultBranchRef struct {
			Target struct {
				History struct {
					TotalCount int `json:"totalCount"`
				} `json:"history"`
			} `json:"target"`
		} `json:"defaultBranchRef"`

		PullRequests struct {
			TotalCount int `json:"totalCount"`
			Nodes      []struct {
				State          string `json:"state"` // OPEN, CLOSED, MERGED
				CreatedAt      string `json:"createdAt"`
				ReviewRequests struct {
					TotalCount int `json:"totalCount"`
				} `json:"reviewRequests"`
			} `json:"nodes"`
		} `json:"pullRequests"`

		Ref struct {
			Target struct {
				History struct {
					Edges []struct {
						Node struct {
							CommittedDate string `json:"committedDate"`
							Additions     int    `json:"additions"`
							Deletions     int    `json:"deletions"`
						} `json:"node"`
					} `json:"edges"`
				} `json:"history"`
			} `json:"target"`
		} `json:"ref"`
	} `json:"repository"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}