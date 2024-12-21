package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lutfifadlan/directories/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/api/directories", handler.AddDirectory)
	r.GET("/api/directories/:id", handler.GetDirectoryById)
	r.POST("/api/users", handler.AddUser)
	r.GET("/api/users/:id", handler.GetUserById)
	r.POST("/api/magic-link", handler.CreateMagicLink)
	r.GET("/", gin.WrapF(handler.IndexHandler))
	r.GET("/ping", handler.Ping)
}
