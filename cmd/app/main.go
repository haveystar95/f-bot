package main

import (
	"f-bot/internal/handlers"
	"f-bot/internal/services"
	"f-bot/pkg/config"
	"f-bot/pkg/db"
	"f-bot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := config.LoadConfig()
	logger.InitLogger()
	db.InitDB(cfg.Database.Host, strconv.Itoa(cfg.Database.Port), cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	r := gin.Default()

	userService := services.NewUserService(db.DB)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users/:id", userHandler.GetUser)

	err = r.Run(":8084")
	if err != nil {
		return
	}
}
