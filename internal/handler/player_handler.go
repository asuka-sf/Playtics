package handler

import (
	"errors"
	"net/http"

	"playtics/internal/domain"
	"playtics/internal/usecase"

	"github.com/gin-gonic/gin"
)

type playerHandler struct {
	us usecase.PlayerUsecase
}

type createPlayerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageURL string `json:"image_url"`
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
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call usecase
	result, err := h.us.Create(ctx, req.Name, req.Email, req.ImageURL)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"player": result})
}
