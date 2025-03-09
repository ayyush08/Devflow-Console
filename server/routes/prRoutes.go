package routes

import (
	"github.com/ayyush08/keploy-dashboard/controllers"
	"github.com/gin-gonic/gin"
)

func PRRoutes(api *gin.RouterGroup){
	metrics := api.Group("/metrics")
	{
		metrics.GET("/pr/:owner/:repo", controllers.GetPRMetrics)
	}
}