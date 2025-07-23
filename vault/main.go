package main

import (
	"net/http"
	"os"
	"time"

	"github.com/elaurentium/vaultui/internal/domain/services"
	"github.com/elaurentium/vaultui/internal/infra/api"
	"github.com/elaurentium/vaultui/internal/infra/api/handlers"
	"github.com/elaurentium/vaultui/internal/infra/api/middlewares"
	"github.com/elaurentium/vaultui/internal/infra/auth"
	"github.com/elaurentium/vaultui/internal/infra/persistence"
	"github.com/elaurentium/vaultui/internal/infra/persistence/db"
	"github.com/elaurentium/vaultui/internal/infra/persistence/redis"
	"github.com/elaurentium/vaultui/internal/utils/log"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger := log.NewLogger()
	logger.Info("Starting Application")

	redisClient, err := redis.NewRedisClient()
	if err != nil {
		logger.Error("Failed to connect to redis: %v", err)
		return
	}
	defer redisClient.Close()

	database, err := persistence.NewSqlite()
	if err != nil {
		logger.Error("Failed to connect with db: %v", err)
		return
	}
	defer database.Close()

	userRepo := db.NewUserRepository(database)
	authService := auth.NewAuthService()
	userService := services.NewUserService(userRepo, authService)
	userHandler := handlers.NewUserHandler(userService)

	authMiddleware := &middlewares.AuthMiddleware{}

	router := api.NewRouter(userHandler, authMiddleware, redisClient)

	server := &http.Server{
		Addr: os.Getenv("ADDR"),
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Failed to start server: %v", err.Error())
	}
}