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

type MockPlayerRepository struct {
	player *domain.Player
	err    error
}

func (m *MockPlayerRepository) Create(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	return m.player, m.err
}

func TestCreatePlayer(t *testing.T) {
	tests := []struct {
		name   string
		args   *domain.Player
		result *domain.Player
		err    error
	}{
		{
			name: "success",
			args: &domain.Player{
				Name:     "Alice",
				Email:    "alice@email.com",
				ImageURL: "test_image.jpg",
			},
			result: &domain.Player{
				ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
				Name:      "Alice",
				Email:     "alice@email.com",
				ImageURL:  "test_image.jpg",
				CreatedAt: time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
			},
			err: nil,
		},
		{
			name: "success: no image_url",
			args: &domain.Player{
				Name:     "Alice",
				Email:    "alice@email.com",
				ImageURL: "",
			},
			result: &domain.Player{
				ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
				Name:      "Alice",
				Email:     "alice@email.com",
				ImageURL:  "",
				CreatedAt: time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
			},
			err: nil,
		},
		{
			name: "failed to create player",
			args: &domain.Player{
				Name:  "Alice",
				Email: "alice@email.com",
			},
			err: errors.New("failed to create player"),
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &MockPlayerRepository{
				player: tt.result,
				err:    tt.err,
			}
			uc := NewPlayerUsecase(mock)
			player, err := uc.Create(ctx, tt.args.Name, tt.args.Email, tt.args.ImageURL)

			// check if the result and error has the values as expected
			if tt.err != nil {
				assert.Error(t, err)
				assert.Nil(t, player)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.result, player)
			}
		})
	}
}
