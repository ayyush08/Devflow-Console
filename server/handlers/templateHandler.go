package handlers

import "github.com/ayyush08/keploy-dashboard/models"

func ApplyTemplate(template string, metrics models.DashboardMetrics) interface{} {

	switch template {
	case "developer":
		{
			return models.DeveloperTemplate{
				Merged:       metrics.PRMetrics.Merged,
				Pending:      metrics.PRMetrics.Pending,
				AvgMergeTime: metrics.PRMetrics.AvgMergeTime,
			}
		}
	case "qa":
		{
			return models.QATemplate{
				Rejected:    metrics.PRMetrics.Rejected,
				TestsPassed: metrics.TestMetrics.TestsPassed,
				TestsFailed: metrics.TestMetrics.TestsFailed,
				PassRate:    metrics.TestMetrics.PassRate,
			}
		}
	case "manager":
		{
			return models.ManagerTemplate{
				PRCount:    metrics.PRMetrics.OpenPRCount + metrics.PRMetrics.ClosedPRCount + metrics.PRMetrics.MergedPRCount,
				Stars:      metrics.RepoMetrics.Stars,
				OpenIssues: metrics.RepoMetrics.OpenIssues,
			}
		}
	default:
		{
			return metrics
		}
	}
}
