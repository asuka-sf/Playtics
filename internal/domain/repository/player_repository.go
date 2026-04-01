package repository

import (
	"context"

	"playtics/internal/domain"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *domain.Player) (*domain.Player, error)
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Player, error)
}
