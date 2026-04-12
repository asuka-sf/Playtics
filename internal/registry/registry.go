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
	matchResultRepo := repository.NewMatchResultRepository(queries)

	// usecase
	playerUsecase := usecase.NewPlayerUsecase(playerRepo)
	matchUsecase := usecase.NewMatchUsecase(matchRepo)
	matchResultUsecase := usecase.NewMatchResultUsecase(matchResultRepo)

	// handler
	playerHandler := handler.NewPlayerHandler(playerUsecase)
	matchHandler := handler.NewMatchHandler(matchUsecase)
	matchResultHandler := handler.NewMatchResultHandler(matchResultUsecase)

	// appHandler
	appHandler := handler.NewHandler(playerHandler, matchHandler, matchResultHandler)

	return &Registry{
		AppHandler: appHandler,
	}
}
