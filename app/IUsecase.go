package app

import "github.com/gorilla/websocket"

type IUsecase interface {
	ClientHandler(conn *websocket.Conn)
}
