package repository

import (
	"context"

	"playtics/internal/domain"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *domain.Player) (*domain.Player, error)
}
