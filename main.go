package main

import (
	"rate-limiter-service/internal/handlers"
	"rate-limiter-service/internal/service"
	"rate-limiter-service/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := store.NewMemoryStore()
	service := service.NewRateLimiterService(store)

	handler := handlers.NewRateLimiterHandler(service)
	r.GET("/allow", handler.AllowRequest)
	r.Run(":8080")
}
