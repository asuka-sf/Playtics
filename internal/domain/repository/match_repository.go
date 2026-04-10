package repository

import (
	"context"

	"playtics/internal/domain"
)

type MatchRepository interface {
	Create(ctx context.Context, match *domain.Match) (*domain.Match, error)
}
