package helpers

import (
  "log"
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)

type Message struct {
	Nickname      string    `json:"nickname"`
	Fullname      string    `json:"fullname" `
	UserId        uint 	    `json:"userId"`
	ChannelId     uint      `json:"channelId"`
  Message       string    `json:"message" valid:"required"`
  MessageId     int       `json:"messageId"`
}


var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)

var socketUpgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
  ReadBufferSize:  1024,
  WriteBufferSize: 1024,
}

func InitSocket() {
  go HandleMessages()
}


func HandleIncoming(w http.ResponseWriter, r *http.Request, isChatBot bool) {
	// Upgrade initial GET request to a websocket
	conn, err := socketUpgrader.Upgrade(w, r, nil)
  if err != nil {
      fmt.Println("Failed to set websocket upgrade: %v", err)
      return
  }
	// Make sure we close the connection when the function returns
	defer conn.Close()

	// Register our new client
	clients[conn] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, conn)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
