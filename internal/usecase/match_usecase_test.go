package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"playtics/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockMatchRepository struct {
	match *domain.Match
	err   error
}

func (m *MockMatchRepository) Create(ctx context.Context, match *domain.Match) (*domain.Match, error) {
	return m.match, m.err
}

func TestCreateMatch(t *testing.T) {
	tests := []struct {
		name   string
		args   *domain.Match
		result *domain.Match
		err    error
	}{
		{
			name: "success",
			args: &domain.Match{
				DurationSeconds: 300,
			},
			result: &domain.Match{
				ID:              uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
				DurationSeconds: 300,
				CreatedAt:       time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
			},
			err: nil,
		},
		{
			name: "failed to create match",
			args: &domain.Match{
				DurationSeconds: 300,
			},
			err: errors.New("failed to create match"),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockMatchRepository{
				match: tt.result,
				err:   tt.err,
			}

			uc := NewMatchUsecase(mock)
			match, err := uc.Create(ctx, tt.args.DurationSeconds)

			if tt.err != nil {
				assert.Error(t, err)
				assert.Nil(t, match)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, match)
			}

		})
	}
}
