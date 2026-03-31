package handler

type Handler struct {
	playerHandler *playerHandler
}

func NewHandler(playerHandler *playerHandler) *Handler {
	return &Handler{
		playerHandler: playerHandler,
	}
}
