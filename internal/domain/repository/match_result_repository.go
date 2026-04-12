package repository

import (
	"context"

	"playtics/internal/domain"
)

type MatchResultRepository interface {
	Create(ctx context.Context, matchResult *domain.MatchResult) (*domain.MatchResult, error)
}
