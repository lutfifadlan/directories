package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lutfifadlan/directories/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/directories", handler.AddDirectory)
	r.GET("/ping", handler.Ping)
}
