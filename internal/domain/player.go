package domain

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID
	Name      string
	Email     string
	ImageURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
