package main

import (
	"log"

	"task-manager/internal/handler"
	"task-manager/internal/repository"
	"task-manager/internal/router"
	"task-manager/internal/service"
	"task-manager/pkg/config"
	"task-manager/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig("config.yaml")

	db.ConnectDatabase(cfg)

	//	ctx := context.Background()
	r := gin.Default()

	// Register routes
	userRepository := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	router.RegisterUserRoutes(r, userHandler)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("failed to run server", err)
	}
}
