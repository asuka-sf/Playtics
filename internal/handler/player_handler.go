package handler

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"playtics/internal/domain"
	"playtics/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

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

type createPlayerRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
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

// map JSON request to createPlayerRequest structure
func bindJSON(c *gin.Context, req *createPlayerRequest) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		var ve validator.ValidationErrors
		// check if the error is validator.ValidationErrors and get the error information
		if errors.As(err, &ve) {
			// get the information of createPlayerRequest
			structType := reflect.TypeOf(createPlayerRequest{})
			var msgs []string
			for _, fieldErr := range ve {
				// check if the field exists in the struct and get the json tag name
				if structField, ok := structType.FieldByName(fieldErr.Field()); ok {
					jsonTag := structField.Tag.Get("json")
					msgs = append(msgs, jsonTag+" is "+fieldErr.Tag())
				}
			}
			// set a message (e.g. "name is required, email is required")
			message := strings.Join(msgs, ", ")
			if message == "" {
				message = "validation failed"
			}
			errorResponse(c, http.StatusBadRequest, message)
			return false
		}
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return false
	}
	return true
}
