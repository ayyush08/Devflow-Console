package controllers

import (
	"net/http"

	"github.com/ayyush08/keploy-dashboard/handlers"
	"github.com/ayyush08/keploy-dashboard/models"
	"github.com/gin-gonic/gin"
)

func GetPRMetrics(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")
	template := c.Param("template")


	if owner == "" || repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner and repo are required"})
		return
	}

	if len(template) > 0 {
		template = template[1:] 
	} else {
		template = "default"
	}


	prMetrics, err :=	 handlers.FetchPRMetrics(owner, repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}



	testMetrics, err := handlers.FetchTestMetrics(owner, repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	repoMetrics, err := handlers.FetchRepoMetrics(owner, repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dashboardMetrics models.DashboardMetrics
	dashboardMetrics.PRMetrics = prMetrics
	dashboardMetrics.TestMetrics = testMetrics
	dashboardMetrics.RepoMetrics = repoMetrics


	dashboard := handlers.ApplyTemplate(template, dashboardMetrics)

	c.JSON(http.StatusOK, dashboard)
}
