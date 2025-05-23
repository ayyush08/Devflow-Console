package controllers

import (
	"net/http"

	"github.com/ayyush08/devflow-console/server/handlers"
	"github.com/gin-gonic/gin"
)



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
