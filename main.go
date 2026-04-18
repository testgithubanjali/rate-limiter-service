package main

import (
	"rate-limiter-service/internal/handlers"
	"rate-limiter-service/internal/services"
	"rate-limiter-service/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := store.NewMemoryStore()
	services := services.NewRateLimiterService(store)

	handler := handlers.NewRateLimiterHandler(services)
	r.GET("/allow", handler.AllowRequest)
	r.Run(":8080")
}
