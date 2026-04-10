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

type createPlayerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	ImageURL string `json:"image_url"`
}

type playerResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type playerHandler struct {
	us usecase.PlayerUsecase
}

func NewPlayerHandler(us usecase.PlayerUsecase) *playerHandler {
	return &playerHandler{
		us: us,
	}
}

func (h *playerHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	// get values from request
	var req createPlayerRequest
	if !bindJSON(c, &req) {
		return
	}

	// call usecase
	result, err := h.us.Create(ctx, req.Name, req.Email, req.ImageURL)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			errorResponse(c, http.StatusConflict, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		log.Printf("failed to create player: %v", err)
		return
	}

	successResponse(c, http.StatusCreated, "Created player successfully", playerResponse{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		ImageURL:  result.ImageURL,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (h *playerHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	idStr := c.Param("id")
	// parse string to uuid.UUID
	id, err := uuid.Parse(idStr)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid player id")
		return
	}

	result, err := h.us.GetByID(ctx, id)
	if err != nil {
		// if can't find the player
		if errors.Is(err, domain.ErrNotFound) {
			errorResponse(c, http.StatusNotFound, "player not found")
			return
		}
		errorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	successResponse(c, http.StatusOK, "Fetched player successfully", playerResponse{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		ImageURL:  result.ImageURL,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}
