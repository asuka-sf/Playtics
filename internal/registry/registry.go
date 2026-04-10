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
	matchRepo := repository.NewMatchRepository(queries)

	// usecase
	playerUsecase := usecase.NewPlayerUsecase(playerRepo)
	matchUsecase := usecase.NewMatchUsecase(matchRepo)

	// handler
	playerHandler := handler.NewPlayerHandler(playerUsecase)
	matchHandler := handler.NewMatchHandler(matchUsecase)

	// appHandler
	appHandler := handler.NewHandler(playerHandler, matchHandler)

	return &Registry{
		AppHandler: appHandler,
	}
}
