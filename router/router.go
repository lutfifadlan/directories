package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lutfifadlan/directories/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.POST("/api/directories", handler.AddDirectory)
	r.GET("/api/directories/:id", handler.GetDirectoryById)
	r.POST("/api/users", handler.AddUser)
	r.GET("/api/users/:id", handler.GetUserById)
	r.POST("/api/magic-links", handler.CreateMagicLink)
	r.GET("/api/magic-links/:token", handler.VerifyMagicLink)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/ping", handler.Ping)
}
