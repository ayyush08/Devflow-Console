package routes

import (
	"github.com/ayyush08/devflow-console/controllers"
	"github.com/gin-gonic/gin"
)

func MetricRoutes(api *gin.RouterGroup){
	metrics := api.Group("/metrics")
	{
		metrics.GET("/:owner/:repo/*template", controllers.GetTemplatizedMetrics)
		metrics.GET("/general/:owner/:repo", controllers.GetGeneralMetrics)
		metrics.GET("/developer/:owner/:repo", controllers.GetDevTemplate)
		metrics.GET("/qa/:owner/:repo", controllers.GetQaTemplate)
	}
}
