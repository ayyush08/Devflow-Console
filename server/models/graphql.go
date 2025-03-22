package models







type GraphQLRequest struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
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
