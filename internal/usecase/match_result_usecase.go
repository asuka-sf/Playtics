package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"playtics/internal/domain"
	"playtics/internal/domain/repository"
)

type MatchResultUsecase interface {
	Create(ctx context.Context, playerID, matchID uuid.UUID, killCount, deathCount, score int) (*domain.MatchResult, error)
}

type matchResultUsecase struct {
	repo repository.MatchResultRepository
}

func NewMatchResultUsecase(repo repository.MatchResultRepository) *matchResultUsecase {
	return &matchResultUsecase{
		repo: repo,
	}
}

func (u *matchResultUsecase) Create(ctx context.Context, playerID, matchID uuid.UUID, killCount, deathCount, score int) (*domain.MatchResult, error) {
	// validation check
	if killCount < 0 {
		return nil, fmt.Errorf("kill_count must be 0 or greater: %w", domain.ErrValidation)
	}
	if deathCount < 0 {
		return nil, fmt.Errorf("death_count must be 0 or greater: %w", domain.ErrValidation)
	}
	if score < 0 {
		return nil, fmt.Errorf("score must be 0 or greater: %w", domain.ErrValidation)
	}

	matchResult, err := u.repo.Create(ctx, &domain.MatchResult{
		PlayerID:   playerID,
		MatchID:    matchID,
		KillCount:  killCount,
		DeathCount: deathCount,
		Score:      score,
	})

	if err != nil {
		return nil, err
	}

	return matchResult, nil
}
