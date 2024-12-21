package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lutfifadlan/directories/internal/repository"
	"github.com/lutfifadlan/directories/internal/service"
)

var magicLinkService *service.MagicLinkService
var userRepository *repository.UserRepository

func SetMagicLinkService(svc *service.MagicLinkService) {
	magicLinkService = svc
}

func SetUserRepository(repo *repository.UserRepository) {
	userRepository = repo
}

func CreateMagicLink(c *gin.Context) {
	var newMagicLink struct {
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&newMagicLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := magicLinkService.GenerateMagicLink(userRepository.DB, newMagicLink.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, m)
}

func VerifyMagicLink(c *gin.Context) {
	token := c.Param("token")

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	m, err := magicLinkService.VerifyMagicLink(userRepository.DB, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if m == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "magic link not found"})
		return
	}

	c.JSON(http.StatusOK, m)
}
