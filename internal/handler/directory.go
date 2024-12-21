package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lutfifadlan/directories/internal/service"
)

var directoryService *service.DirectoryService

func SetDirectoryService(svc *service.DirectoryService) {
	directoryService = svc
}

func AddDirectory(c *gin.Context) {
	var newDirectory struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&newDirectory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := directoryService.AddDirectory(newDirectory.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, d)
}
