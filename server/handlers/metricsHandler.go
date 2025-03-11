package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ayyush08/keploy-dashboard/config"
	"github.com/ayyush08/keploy-dashboard/models"
)

const query = `
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
						checkSuites(first: 10) {
							nodes {
								conclusion
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
	}
`

func FetchMetrics(owner string, repo string) (models.DashboardMetrics, error) {

	graphQLPayload := models.GraphQLRequest{
		Query: query,
		Variables: map[string]string{
			"owner": owner,
			"name":  repo,
		},
	}

	jsonPayload, err := json.Marshal(graphQLPayload)

	if err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("error marshalling graphql payload: %v", err)
	}

	req, _ := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+config.GetGithubToken())
	req.Header.Set("Content-Type", "application/json")


	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("failed to fetch metrics: %v", err)
	}

	defer res.Body.Close()
	
	data,err := io.ReadAll(res.Body)

	if err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var graphQLResponse models.GraphQLResponse


	if err := json.Unmarshal(data, &graphQLResponse); err != nil {
		return models.DashboardMetrics{}, fmt.Errorf("failed to parse GraphQL response: %v", err)
	}

	if len(graphQLResponse.Errors) > 0 {
		return models.DashboardMetrics{}, fmt.Errorf("GraphQL error: %v", graphQLResponse.Errors[0].Message)
	}


	var transformedMetrics models.DashboardMetrics = extractMetrics(graphQLResponse)



	return transformedMetrics, nil
	
}


func extractMetrics(graphQLResponse models.GraphQLResponse) models.DashboardMetrics {
	

	return models.DashboardMetrics{}
}