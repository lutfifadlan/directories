package main

import (
	"log"

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

	directoryService := service.NewDirectoryService(directoryRepo)
	userService := service.NewUserService(userRepo)

	handler.SetDirectoryService(directoryService)
	handler.SetUserService(userService)

	r := gin.Default()
	r.Use(gin.Recovery())

	router.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed: ", err)
	}
}
