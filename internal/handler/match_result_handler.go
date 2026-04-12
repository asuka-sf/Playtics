package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"playtics/internal/domain"
	"playtics/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createMatchResultRequest struct {
	// Note: PlayerID is string (not uuid.UUID) because binding:"required" does not
	// reject omitted fixed-length array types like uuid.UUID. Parse to uuid in handler.
	PlayerID   string `json:"player_id" binding:"required"`
	KillCount  int    `json:"kill_count"`
	DeathCount int    `json:"death_count"`
	Score      int    `json:"score"`
}

type matchResultResponse struct {
	PlayerID   uuid.UUID `json:"player_id"`
	MatchID    uuid.UUID `json:"match_id"`
	KillCount  int       `json:"kill_count"`
	DeathCount int       `json:"death_count"`
	Score      int       `json:"score"`
	CreatedAt  time.Time `json:"created_at"`
}

type matchResultHandler struct {
	uc usecase.MatchResultUsecase
}

func NewMatchResultHandler(uc usecase.MatchResultUsecase) *matchResultHandler {
	return &matchResultHandler{
		uc: uc,
	}
}

func (h *matchResultHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	// get match id from url path
	matchIDStr := c.Param("id")
	matchID, err := uuid.Parse(matchIDStr)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid match id")
		return
	}

	var req createMatchResultRequest
	if !bindJSON(c, &req) {
		return
	}

	// parse player id string to uuid.UUID
	playerID, err := uuid.Parse(req.PlayerID)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid player id")
		return
	}

	result, err := h.uc.Create(ctx, playerID, matchID, req.KillCount, req.DeathCount, req.Score)
	if err != nil {
		// if player_id or match_id doesn't exist
		if errors.Is(err, domain.ErrNotFound) {
			errorResponse(c, http.StatusNotFound, "player or match not found")
			return
		}
		if errors.Is(err, domain.ErrValidation) {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		log.Printf("failed to create match result: %v", err)
		return
	}

	successResponse(c, http.StatusCreated, "Created match result successfully", matchResultResponse{
		PlayerID:   result.PlayerID,
		MatchID:    result.MatchID,
		KillCount:  result.KillCount,
		DeathCount: result.DeathCount,
		Score:      result.Score,
		CreatedAt:  result.CreatedAt,
	})
}
