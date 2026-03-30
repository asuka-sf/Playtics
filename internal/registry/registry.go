package registry

import (
	"playtics/internal/handler"
	"playtics/internal/infrastructure/postgres/gen"
	"playtics/internal/infrastructure/postgres/repository"
	"playtics/internal/usecase"
)

type Registry struct {
	AppHandler *handler.Handler
}

func NewRegistry(queries *gen.Queries) *Registry {

	playerRepo := repository.NewPlayerRepository(queries)

	// usecase
	playerUsecase := usecase.NewPlayerUsecase(playerRepo)

	// handler
	playerHandler := handler.NewPlayerHandler(playerUsecase)

	// appHandler
	appHandler := handler.NewHandler(playerHandler)

	return &Registry{
		AppHandler: appHandler,
	}
}
