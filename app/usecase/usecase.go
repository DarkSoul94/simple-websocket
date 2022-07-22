package usecase

import (
	"encoding/json"

	"github.com/DarkSoul94/simple-websocket/app"
	"github.com/DarkSoul94/simple-websocket/models"
	"github.com/gorilla/websocket"
)

type usecase struct{}

func NewUsecase() app.IUsecase {
	return &usecase{}
}

func (u *usecase) ClientHandler(conn *websocket.Conn) {
	defer conn.Close()

	var req models.Message

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}

		err = json.Unmarshal(message, &req)
		if err != nil {
			break
		}

		switch req.Event {
		case "ping":
			resp := models.Message{
				Event: "pong",
			}

			err = conn.WriteJSON(resp)
			if err != nil {

			}
		}
	}
}
