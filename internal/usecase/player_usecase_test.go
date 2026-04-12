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

func (m *MockPlayerRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Player, error) {
	return m.player, m.err
}

var wantPlayer = &domain.Player{
	ID:        uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
	Name:      "Alice",
	Email:     "alice@email.com",
	ImageURL:  "test_image.jpg",
	CreatedAt: time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2026, 3, 30, 10, 0, 0, 0, time.UTC),
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
			result: wantPlayer,
			err:    nil,
		},
		{
			name: "success: no image_url",
			args: &domain.Player{
				Name:     "Alice",
				Email:    "alice@email.com",
				ImageURL: "",
			},
			result: wantPlayer,
			err:    nil,
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

func TestGetPlayer(t *testing.T) {
	tests := []struct {
		name   string
		result *domain.Player
		err    error
	}{
		{
			name:   "success",
			result: wantPlayer,
			err:    nil,
		},
		{
			name: "player not found",
			err:  domain.ErrNotFound,
		},
		{
			name: "internal server error",
			err:  errors.New("connection refused"),
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
			id := uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")
			player, err := uc.GetByID(ctx, id)

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
