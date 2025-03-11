package models







type GraphQLRequest struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
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
	}
}
