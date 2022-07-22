package http

import (
	"encoding/json"
	"net/http"

	"github.com/DarkSoul94/simple-websocket/app"
	"github.com/DarkSoul94/simple-websocket/models"
	"github.com/gorilla/websocket"
)

// Handler ...
type Handler struct {
	uc app.IUsecase
}

// NewHandler ...
func NewHandler(uc app.IUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "error"})
	}

	for {
		var req models.Message

		mt, message, err := ws.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}

		err = json.Unmarshal(message, &req)
		if err != nil {
			break
		}

		resp := h.uc.MessageHandler(req)

		err = ws.WriteJSON(resp)
		if err != nil {
			break
		}
	}
}
