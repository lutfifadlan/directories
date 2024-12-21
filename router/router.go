package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lutfifadlan/directories/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/directories", handler.AddDirectory)
	r.GET("/directories/:id", handler.GetDirectoryById)
	r.POST("/users", handler.AddUser)
	r.GET("/users/:id", handler.GetUserById)
	r.GET("/ping", handler.Ping)
}
