package domain

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID              uuid.UUID
	DurationSeconds int
	CreatedAt       time.Time
}
