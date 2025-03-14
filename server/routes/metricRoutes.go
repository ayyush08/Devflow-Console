package routes

import (
	"github.com/ayyush08/devflow-console/controllers"
	"github.com/gin-gonic/gin"
)

func PRRoutes(api *gin.RouterGroup){
	metrics := api.Group("/metrics")
	{
		metrics.GET("/:owner/:repo/*template", controllers.GetPRMetrics)
	}
}
