package api

import (
	"github.com/elaurentium/vaultui/internal/infra/api/handlers"
	"github.com/elaurentium/vaultui/internal/infra/api/middlewares"
	"github.com/elaurentium/vaultui/internal/infra/persistence/redis"
	"github.com/elaurentium/vaultui/internal/utils/log"
	"github.com/gin-gonic/gin"
)

func NewRouter(UserHandler *handlers.UserHandler, authMiddleware *middlewares.AuthMiddleware, redisClient *redis.RedisClient) *gin.Engine {
	log := log.NewLogger()
	router := gin.Default()

	router.Use(middlewares.CorsMiddleware())
	router.Use(middlewares.LoggerMiddleware(log))
	router.Use(middlewares.SecurityMiddleware())
	router.Use(middlewares.NewRateLimiterMiddleware(redisClient.GetClient()))

	router.POST("/register", UserHandler.Register)
	router.POST("/login", UserHandler.Login)

	return router
}
