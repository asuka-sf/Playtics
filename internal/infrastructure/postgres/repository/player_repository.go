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

type playerRepository struct {
	queries *gen.Queries
}

func NewPlayerRepository(queries *gen.Queries) *playerRepository {
	return &playerRepository{
		queries: queries,
	}
}

// create player
func (r *playerRepository) Create(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	result, err := r.queries.CreatePlayer(ctx, gen.CreatePlayerParams{
		ID:       pgtype.UUID{Bytes: player.ID, Valid: true},
		Name:     player.Name,
		Email:    player.Email,
		ImageUrl: pgtype.Text{String: player.ImageURL, Valid: player.ImageURL != ""},
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, domain.ErrDuplicateEmail
		}
		return nil, fmt.Errorf("failed to create player: %w", err)
	}

	return &domain.Player{
		ID:        result.ID.Bytes,
		Name:      result.Name,
		Email:     result.Email,
		ImageURL:  result.ImageUrl.String,
		CreatedAt: result.CreatedAt.Time,
		UpdatedAt: result.UpdatedAt.Time,
	}, nil
}
