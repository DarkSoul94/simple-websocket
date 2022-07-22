package http

import (
	"net/http"

	"github.com/DarkSoul94/simple-websocket/app"
)

// RegisterHTTPEndpoints ...
func RegisterHTTPEndpoints(mux *http.ServeMux, uc app.IUsecase) {
	h := NewHandler(uc)

	mux.HandleFunc("/ws", h.WsEndpoint)
}
