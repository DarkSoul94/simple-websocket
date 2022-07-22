package usecase

import (
	"github.com/DarkSoul94/simple-websocket/app"
	"github.com/DarkSoul94/simple-websocket/models"
)

type usecase struct{}

func NewUsecase() app.IUsecase {
	return &usecase{}
}

func (u *usecase) MessageHandler(req models.Message) models.Message {
	switch req.Event {
	case "ping":
		resp := models.Message{
			Event: "pong",
		}
		return resp
	default:
		return models.Message{}
	}
}
