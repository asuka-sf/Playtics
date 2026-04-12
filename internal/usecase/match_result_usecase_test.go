package usecase

import (
	"context"
	"errors"
	"playtics/internal/domain"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var matchResult = &domain.MatchResult{
	PlayerID:   uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
	MatchID:    uuid.MustParse("550e8400-e29b-41d4-a716-446655770000"),
	KillCount:  100,
	DeathCount: 80,
	Score:      10000,
	CreatedAt:  time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
}

type MockMatchResultRepository struct {
	matchResult *domain.MatchResult
	err         error
}

func (m *MockMatchResultRepository) Create(ctx context.Context, matchResult *domain.MatchResult) (*domain.MatchResult, error) {
	return m.matchResult, m.err
}

func TestMatchResultCreate(t *testing.T) {
	tests := []struct {
		name    string
		args    *domain.MatchResult
		result  *domain.MatchResult
		mockErr error
		wantErr bool
	}{
		{
			name:    "success",
			args:    matchResult,
			result:  matchResult,
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "player or match doesn't exist",
			args:    &domain.MatchResult{},
			mockErr: domain.ErrNotFound,
			wantErr: true,
		},
		{
			name: "invalid kill_count",
			args: &domain.MatchResult{
				PlayerID:   uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
				MatchID:    uuid.MustParse("550e8400-e29b-41d4-a716-446655770000"),
				KillCount:  -1,
				DeathCount: 80,
				Score:      10000,
				CreatedAt:  time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
			},
			mockErr: nil,
			wantErr: true,
		},
		{
			name:    "failed to create match result",
			args:    matchResult,
			mockErr: errors.New("failed to create match result"),
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockMatchResultRepository{
				matchResult: tt.result,
				err:         tt.mockErr,
			}

			uc := NewMatchResultUsecase(mock)
			matchResult, err := uc.Create(ctx, tt.args.PlayerID, tt.args.MatchID, tt.args.KillCount, tt.args.DeathCount, tt.args.Score)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, matchResult)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, matchResult)
			}
		})
	}
}
