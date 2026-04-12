package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"playtics/internal/domain"
	"playtics/internal/domain/repository"
)

type MatchResultUsecase interface {
	Create(ctx context.Context, playerId, matchId uuid.UUID, killCount, deathCount, score int) (*domain.MatchResult, error)
}

type matchResultUsecase struct {
	repo repository.MatchResultRepository
}

func NewMatchResultUsecase(repo repository.MatchResultRepository) *matchResultUsecase {
	return &matchResultUsecase{
		repo: repo,
	}
}

func (u *matchResultUsecase) Create(ctx context.Context, playerId, matchId uuid.UUID, killCount, deathCount, score int) (*domain.MatchResult, error) {
	// validation check
	if killCount < 0 {
		return nil, errors.New("kill_count must be 0 or greater")
	}
	if deathCount < 0 {
		return nil, errors.New("death_count must be 0 or greater")
	}
	if score < 0 {
		return nil, errors.New("score must be 0 or greater")
	}

	matchResult, err := u.repo.Create(ctx, &domain.MatchResult{
		PlayerID:   playerId,
		MatchID:    matchId,
		KillCount:  killCount,
		DeathCount: deathCount,
		Score:      score,
	})

	if err != nil {
		return nil, err
	}

	return matchResult, nil
}
