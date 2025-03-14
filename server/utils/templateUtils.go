package utils

import "github.com/ayyush08/devflow-console/models"

func ApplyQaTemplate(metrics models.DashboardMetrics) models.QATemplate {

	var qaResponse models.QATemplate

	qaResponse.Conclusions = make(map[models.TestConclusion]int)

	var passedTests, failedTest, totalTests int

	for conclusion, count := range metrics.TestMetrics.Conclusions {
		qaResponse.Conclusions[conclusion] = count 

		if conclusion == models.ConclusionSuccess {
			passedTests += count
		} else if conclusion == models.ConclusionFailure || conclusion == models.ConclusionCanceled {
			failedTest += count
			qaResponse.Rejected += count
		}

		totalTests += count
	}

	qaResponse.TestSuites = append(qaResponse.TestSuites, metrics.TestMetrics.TestSuites...)


	qaResponse.PassRate = float64(passedTests) / float64(totalTests)

	qaResponse.Rejected = failedTest

	return qaResponse

}



func ApplyDeveloperTemplate(metrics models.DashboardMetrics) models.DeveloperTemplate{
	var devResponse models.DeveloperTemplate

	devResponse.Merged = metrics.PRMetrics.MergedPRCount
	devResponse.Pending = metrics.PRMetrics.ClosedPRCount
	devResponse.AvgMergeTime = metrics.PRMetrics.AvgMergeTime

	return devResponse

}


func ApplyManagerTemplate(metrics models.DashboardMetrics) models.ManagerTemplate{
	var managerResponse models.ManagerTemplate

	managerResponse.PRCount = metrics.PRMetrics.OpenPRCount + metrics.PRMetrics.ClosedPRCount + metrics.PRMetrics.MergedPRCount
	managerResponse.Stars = metrics.RepoMetrics.Stars
	managerResponse.OpenIssues = metrics.RepoMetrics.OpenIssues
	managerResponse.ClosedIssues = metrics.RepoMetrics.ClosedIssues

	return managerResponse
}
