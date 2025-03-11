package handlers

import (
	"github.com/ayyush08/keploy-dashboard/models"
	"github.com/ayyush08/keploy-dashboard/utils"
)

func ApplyTemplate(template string, metrics models.DashboardMetrics) interface{} {

	switch template {
	case "developer":
		{
			return utils.ApplyDeveloperTemplate(metrics)
		}
	case "qa":
		{
			return utils.ApplyQaTemplate(metrics)
		}
	case "manager":
		{
			return utils.ApplyManagerTemplate(metrics)
		}
	default:
		{
			return metrics
		}
	}
}
