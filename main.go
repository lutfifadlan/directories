package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lutfifadlan/directories/internal/db"
	"github.com/lutfifadlan/directories/internal/handler"
	"github.com/lutfifadlan/directories/internal/repository"
	"github.com/lutfifadlan/directories/internal/service"
	"github.com/lutfifadlan/directories/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn := db.InitDB()

	directoryRepo := repository.NewDirectoryRepository(dbConn)
	userRepo := repository.NewUserRepository(dbConn)
	magicLinkRepo := repository.NewMagicLinkRepository(dbConn)

	directoryService := service.NewDirectoryService(directoryRepo)
	userService := service.NewUserService(userRepo)
	magicLinkService := service.NewMagicLinkService(magicLinkRepo)

	handler.SetDirectoryService(directoryService)
	handler.SetUserService(userService)
	handler.SetMagicLinkService(magicLinkService)
	handler.SetUserRepository(userRepo)

	r := gin.Default()

	// Add CORS middleware with all origins allowed
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "HX-Request", "HX-Trigger", "HX-Target", "HX-Current-URL"},
		ExposeHeaders:    []string{"Content-Length", "HX-Location", "HX-Trigger", "HX-Redirect"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(gin.Recovery())

	router.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed: ", err)
	}
}
