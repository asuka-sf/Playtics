package handler

type Handler struct {
	PlayerHandler *playerHandler
}

func NewHandler(playerHandler *playerHandler) *Handler {
	return &Handler{
		PlayerHandler: playerHandler,
	}
}
