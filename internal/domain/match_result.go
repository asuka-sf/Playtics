package domain

import (
	"time"

	"github.com/google/uuid"
)

type MatchResult struct {
	PlayerID   uuid.UUID
	MatchID    uuid.UUID
	KillCount  int
	DeathCount int
	Score      int
	CreatedAt  time.Time
}
