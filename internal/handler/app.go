package handler

type Handler struct {
	playerHandler *playerHandler
	matchHandler  *matchHandler
}

func NewHandler(playerHandler *playerHandler, matchHandler *matchHandler) *Handler {
	return &Handler{
		playerHandler: playerHandler,
		matchHandler:  matchHandler,
	}
}
