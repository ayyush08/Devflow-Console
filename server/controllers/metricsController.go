package controllers

import (
	"net/http"

	"github.com/ayyush08/devflow-console/handlers"
	"github.com/gin-gonic/gin"
)

func GetTemplatizedMetrics(c *gin.Context) {
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


	metrics,err := handlers.FetchMetrics(owner, repo, template)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	dashboard := handlers.ApplyTemplate(template, metrics)

	c.JSON(http.StatusOK, dashboard)
}

func GetGeneralMetrics(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")

	if owner == "" || repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "owner and repo are required"})
		return
	}

	generalMetrics,err := handlers.FetchGeneralMetrics(owner, repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, generalMetrics)
}
