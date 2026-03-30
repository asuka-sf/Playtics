package usecase

import (
	"context"
	"fmt"

	"playtics/internal/domain"
	"playtics/internal/domain/repository"

	"github.com/google/uuid"
)

type PlayerUsecase interface {
	Create(ctx context.Context, name, email, imageURL string) (*domain.Player, error)
}

type playerUsecase struct {
	repo repository.PlayerRepository
}

func NewPlayerUsecase(repo repository.PlayerRepository) *playerUsecase {
	return &playerUsecase{
		repo: repo,
	}
}

// create player
func (u *playerUsecase) Create(ctx context.Context, name, email, imageURL string) (*domain.Player, error) {
	player, err := u.repo.Create(ctx, &domain.Player{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		ImageURL: imageURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create player: %w", err)
	}

	return player, nil
}
