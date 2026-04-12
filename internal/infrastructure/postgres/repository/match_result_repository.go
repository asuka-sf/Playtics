package repository

import (
	"context"
	"errors"
	"fmt"

	"playtics/internal/domain"

	"playtics/internal/infrastructure/postgres/gen"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type matchResultRepository struct {
	queries *gen.Queries
}

func NewMatchResultRepository(queries *gen.Queries) *matchResultRepository {
	return &matchResultRepository{
		queries: queries,
	}
}

func (r *matchResultRepository) Create(ctx context.Context, matchResult *domain.MatchResult) (*domain.MatchResult, error) {
	result, err := r.queries.CreateMatchResult(ctx, gen.CreateMatchResultParams{
		PlayerID:   pgtype.UUID{Bytes: matchResult.PlayerID, Valid: true},
		MatchID:    pgtype.UUID{Bytes: matchResult.MatchID, Valid: true},
		KillCount:  int32(matchResult.KillCount),
		DeathCount: int32(matchResult.DeathCount),
		Score:      int32(matchResult.Score),
	})

	if err != nil {
		// if player_id or match_id doesn't exist (foreign key violation)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("failed to create match result: %w", err)
	}

	return &domain.MatchResult{
		PlayerID:   result.PlayerID.Bytes,
		MatchID:    result.MatchID.Bytes,
		KillCount:  int(result.KillCount),
		DeathCount: int(result.DeathCount),
		Score:      int(result.Score),
		CreatedAt:  result.CreatedAt.Time,
	}, nil
}
