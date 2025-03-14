package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ayyush08/devflow-console/config"
	"github.com/ayyush08/devflow-console/models"
	"github.com/ayyush08/devflow-console/utils"
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
	}
`

func FetchMetrics(owner string, repo string) (models.DashboardMetrics, error) {

	cacheKey := fmt.Sprintf("%s/%s", owner, repo)

	if cachedMetrics, found := config.MetricsCache.Get(cacheKey); found {
		return cachedMetrics.(models.DashboardMetrics), nil
	}

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

	data, err := io.ReadAll(res.Body)

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

	var metrics models.DashboardMetrics

	metrics.PRMetrics = utils.ExtractPRMetrics(graphQLResponse)
	metrics.RepoMetrics = utils.ExtractRepoMetrics(graphQLResponse)
	metrics.TestMetrics = utils.ExtractTestMetrics(graphQLResponse)

	return metrics, nil

}
