package utils

import (
	"time"

	"github.com/ayyush08/keploy-dashboard/models"
)

func ExtractPRMetrics(graphQLResponse models.GraphQLResponse) models.PRMetrics {

	repo := graphQLResponse.Data.Repository
	var prMetrics models.PRMetrics

	var openPRCount = repo.OpenPRs.TotalCount
	var closedPRCount = repo.ClosedPRs.TotalCount
	var mergedPRCount = repo.MergedPRs.TotalCount

	prMetrics.OpenPRCount = openPRCount
	prMetrics.ClosedPRCount = closedPRCount
	prMetrics.MergedPRCount = mergedPRCount

	var totalMergeTime = time.Duration(0)

	for _, pr := range repo.MergedPRs.Nodes {

		createdAt, createdAtErr := time.Parse(time.RFC3339, pr.CreatedAt)
		mergedAt, mergedAtErr := time.Parse(time.RFC3339, pr.MergedAt)

		if createdAtErr == nil && mergedAtErr == nil {
			totalMergeTime += mergedAt.Sub(createdAt)
		}

		avgMergeTime := time.Duration(0)

		if mergedPRCount > 0 {
			avgMergeTime = totalMergeTime / time.Duration(mergedPRCount)
		}

		prMetrics.AvgMergeTime = float64(avgMergeTime.Seconds())

	}

	return prMetrics
}

func ExtractRepoMetrics(graphQLResponse models.GraphQLResponse) models.RepoMetrics {

	repo := graphQLResponse.Data.Repository
	var repoMetrics models.RepoMetrics

	repoMetrics.Stars = repo.StargazerCount
	repoMetrics.Forks = repo.ForkCount
	repoMetrics.OpenIssues = repo.OpenIssues.TotalCount
	repoMetrics.ClosedIssues = repo.ClosedIssues.TotalCount
	repoMetrics.LastUpdated = repo.UpdatedAt
	repoMetrics.CreatedAt = repo.CreatedAt

	return repoMetrics

}

func ExtractTestMetrics(graphQLResponse models.GraphQLResponse) models.TestMetrics {

	testNodes := graphQLResponse.Data.Repository.DefaultBranchRef.Target.CheckSuites.Nodes

	var testMetrics models.TestMetrics

	testMetrics.Conclusions = make(map[models.TestConclusion]int)
	testMetrics.Statuses = make(map[models.TestStatus]int)

	testMetrics.TotalTests = len(testNodes)

	for _, testSuite := range testNodes {

		if testSuite.Status == nil || testSuite.Conclusion == nil {
			continue
		}

		testStatus := models.TestStatus(*testSuite.Status)
		testConclusion := models.TestConclusion(*testSuite.Conclusion)

		testMetrics.Statuses[testStatus]++
		testMetrics.Conclusions[testConclusion]++

		workflow := testSuite.WorkflowRun

		if workflow != nil {
			testMetrics.TestSuites = append(testMetrics.TestSuites, models.TestSuite{
				Conclusion: testConclusion,
				Status:     testStatus,
				Workflow:   workflow.Workflow.Name,
			})
		}
	}

	return testMetrics

}
