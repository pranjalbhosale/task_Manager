package main

import (
	"log"

	"task-manager/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig("config.yaml")

	//	ctx := context.Background()
	r := gin.Default()

	// Register routes

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("failed to run server", err)
	}
}
