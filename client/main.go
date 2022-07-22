package main

import (
	"fmt"

	"github.com/DarkSoul94/simple-websocket/models"
	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8888/ws", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	req := models.Message{
		Event: "ping",
	}

	err = conn.WriteJSON(req)
	if err != nil {
		fmt.Println(err)
	}

	var resp models.Message
	err = conn.ReadJSON(&resp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
