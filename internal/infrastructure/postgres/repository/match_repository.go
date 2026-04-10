package repository

import (
	"context"
	"fmt"

	"playtics/internal/domain"
	"playtics/internal/infrastructure/postgres/gen"

	"github.com/jackc/pgx/v5/pgtype"
)

type matchRepository struct {
	queries *gen.Queries
}

func NewMatchRepository(queries *gen.Queries) *matchRepository {
	return &matchRepository{
		queries: queries,
	}
}

func (r *matchRepository) Create(ctx context.Context, match *domain.Match) (*domain.Match, error) {
	result, err := r.queries.CreateMatch(ctx, gen.CreateMatchParams{
		ID:              pgtype.UUID{Bytes: match.ID, Valid: true},
		DurationSeconds: int32(match.DurationSeconds),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create match: %w", err)
	}

	return &domain.Match{
		ID:              result.ID.Bytes,
		DurationSeconds: int(result.DurationSeconds),
		CreatedAt:       result.CreatedAt.Time,
	}, nil
}
