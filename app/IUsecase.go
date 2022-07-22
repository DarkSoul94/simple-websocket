package app

import "github.com/DarkSoul94/simple-websocket/models"

type IUsecase interface {
	MessageHandler(req models.Message) models.Message
}
