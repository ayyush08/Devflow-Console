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

type GeneralMetricsGraphQLResponse struct {
    Data struct {
        Repository struct {
            StargazerCount int `json:"stargazerCount"`
            Issues         struct {
                TotalCount int `json:"totalCount"`
            } `json:"issues"`
            PullRequests struct {
                TotalCount int `json:"totalCount"`
                Edges      []struct {
                    Node struct {
                        CreatedAt string `json:"createdAt"`
                    } `json:"node"`
                } `json:"edges"`
            } `json:"pullRequests"`
			TotalCommits struct {
				Target struct {
					History struct {
						TotalCount int `json:"totalCount"`
					} `json:"history"`
				} `json:"target"`
			} `json:"totalCommits"`
            RecentCommits struct {
                Target struct {
                    History struct {
                        TotalCount int `json:"totalCount"`
                        Edges      []struct {
                            Node struct {
                                CommittedDate string `json:"committedDate"`
                            } `json:"node"`
                        } `json:"edges"`
                    } `json:"history"`
                } `json:"target"`
            } `json:"recentCommits"`
            MergedPRs struct {
                TotalCount int `json:"totalCount"`
            } `json:"mergedPRs"`
            ClosedPRs struct {
                TotalCount int `json:"totalCount"`
            } `json:"closedPRs"`
            OpenPRs struct {
                TotalCount int `json:"totalCount"`
            } `json:"openPRs"`
            BarData struct {
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
            } `json:"barData"`
        } `json:"repository"`
    } `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

type GraphQLResponse struct {
	Data struct {
		Repository struct {
			StargazerCount int `json:"stargazerCount"`
			ForkCount      int `json:"forkCount"`
			OpenIssues     struct {
				TotalCount int `json:"totalCount"`
			} `json:"openIssues"`
			ClosedIssues struct {
				TotalCount int `json:"totalCount"`
			} `json:"closedIssues"`
			UpdatedAt string `json:"updatedAt"`
			CreatedAt string `json:"createdAt"`
			OpenPRs   struct {
				TotalCount int
				Nodes      []struct {
					Title     string `json:"title"`
					CreatedAt string `json:"createdAt"`
				}
			} `json:"openPRs"`
			ClosedPRs struct {
				TotalCount int `json:"totalCount"`
				Nodes      []struct {
					Title     string `json:"title"`
					CreatedAt string `json:"createdAt"`
					ClosedAt  string `json:"closedAt"`
				} `json:"nodes"`
			} `json:"closedPRs"`
			MergedPRs struct {
				TotalCount int `json:"totalCount"`
				Nodes      []struct {
					Title     string `json:"title"`
					CreatedAt string `json:"createdAt"`
					MergedAt  string `json:"mergedAt"`
				} `json:"nodes"`
			} `json:"mergedPRs"`
			DefaultBranchRef struct {
				Target struct {
					CheckSuites struct {
						Nodes []struct {
							Conclusion  *string `json:"conclusion"` 
							Status 	*string `json:"status"`
							WorkflowRun *struct {
								Workflow *struct {
									Name string `json:"name"`
								} `json:"workflow"`
							} `json:"workflowRun"`
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
