package controllers

import (
	"net/http"

	"github.com/ayyush08/keploy-dashboard/handlers"
	"github.com/gin-gonic/gin"
)

func GetPRMetrics(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")

	metrics, err := handlers.FetchPRMetrics(owner, repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metrics)
}
