package handler

import (
	"log"
	"net/http"
	"playtics/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createMatchRequest struct {
	DurationSeconds int `json:"duration_seconds" binding:"required"`
}

type matchResponse struct {
	ID              uuid.UUID `json:"id"`
	DurationSeconds int       `json:"duration_seconds"`
	CreatedAt       time.Time `json:"created_at"`
}

type matchHandler struct {
	uc usecase.MatchUsecase
}

func NewMatchHandler(uc usecase.MatchUsecase) *matchHandler {
	return &matchHandler{
		uc: uc,
	}
}

func (h *matchHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req createMatchRequest
	if !bindJSON(c, &req) {
		return
	}

	result, err := h.uc.Create(ctx, req.DurationSeconds)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		log.Printf("failed to create match: %v", err)
		return
	}

	successResponse(c, http.StatusCreated, "Created match successfully", matchResponse{
		ID:              result.ID,
		DurationSeconds: result.DurationSeconds,
		CreatedAt:       result.CreatedAt,
	})

}
