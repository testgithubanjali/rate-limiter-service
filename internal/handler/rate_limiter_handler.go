package handler

import (
	"net/http"
	"rate-limiter-service/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.RateLimiterService
}

func newHnandler(service *services.RateLimiterService) *Handler {
	return &Handler{service: s}
}
func (h *Handler) AllowHandler(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	allowed := h.service.Allow(userID)
	c.JSON(http.StatusOK, gin.H{"allowed" : true})
	c.JSON(http.StatusTooManyRequests, gin.H{"allowed": false})


}
