package handler

type Handler struct {
	playerHandler      *playerHandler
	matchHandler       *matchHandler
	matchResultHandler *matchResultHandler
}

func NewHandler(playerHandler *playerHandler, matchHandler *matchHandler, matchResultHandler *matchResultHandler) *Handler {
	return &Handler{
		playerHandler:      playerHandler,
		matchHandler:       matchHandler,
		matchResultHandler: matchResultHandler,
	}
}
