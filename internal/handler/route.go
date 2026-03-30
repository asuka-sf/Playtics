package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/players", h.PlayerHandler.Create)
}
