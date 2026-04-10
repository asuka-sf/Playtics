package usecase

import (
	"context"

	"playtics/internal/domain"
	"playtics/internal/domain/repository"

	"github.com/google/uuid"
)

type MatchUsecase interface {
	Create(ctx context.Context, durationSeconds int) (*domain.Match, error)
}

type matchUsecase struct {
	repo repository.MatchRepository
}

func NewMatchUsecase(repo repository.MatchRepository) *matchUsecase {
	return &matchUsecase{
		repo: repo,
	}
}

func (u *matchUsecase) Create(ctx context.Context, durationSeconds int) (*domain.Match, error) {
	match, err := u.repo.Create(ctx, &domain.Match{
		ID:              uuid.New(),
		DurationSeconds: durationSeconds,
	})
	if err != nil {
		return nil, err
	}

	return match, nil
}
