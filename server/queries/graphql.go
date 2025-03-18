package queries



const GenerateMetricsQuery = `
query GetAllCommits($repo: String!, $owner: String!) {
  repository(name: $repo, owner: $owner) {
    stargazerCount
    forkCount
    openIssues: issues(states: OPEN) {
      totalCount
    }
    closedIssues: issues(states: CLOSED) {
      totalCount
    }
    latestRelease{
      tagName
      createdAt
    }
    
    pullRequests{
      totalCount
    }
    refs(first: 1, refPrefix: "refs/heads/") {
      edges {
        node {
          target {
            ... on Commit {
              committedDate
            }
          }
        }
      }
    }
    defaultBranchRef {
      	target {
        ... on Commit {
          	history(first: 100) { 
           totalCount
            edges {
              node {
                committedDate
                
              	}
              
            }
          	}
        }
      }
    }
  }
}
`



const MetricsQuery = `
	query Metrics($owner: String!, $name: String!) {
		repository(owner: $owner, name: $name) {
			stargazerCount
			forkCount
			openIssues: issues(states: OPEN) {
				totalCount
			}
			closedIssues: issues(states: CLOSED) {
				totalCount
			}
			updatedAt
			createdAt

			openPRs: pullRequests(states: OPEN, first: 100) {
				totalCount
				nodes{
					title
					createdAt
				}
			}
			closedPRs: pullRequests(states: CLOSED, first: 100) {
				totalCount
				nodes {
					title
					createdAt
					closedAt
				}
			}
			mergedPRs: pullRequests(states: MERGED, first: 100) {
				totalCount
				nodes {
					title
					createdAt
					mergedAt
				}
			}

			defaultBranchRef {
				target {
					... on Commit {
						checkSuites(first: 100) {
							nodes {
								conclusion
								status
								workflowRun {
									workflow {
										name
									}
								}
							}
						}
					}
				}
			}
		}
	}`