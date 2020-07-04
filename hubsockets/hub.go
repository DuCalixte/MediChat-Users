package hubsockets

import(
  "fmt"
  "time"
)
// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

  // Adding a Channel ID to the HUB
  channelId uint
}

type Message struct {
	Nickname      string    `json:"nickname"`
	Fullname      string    `json:"fullname" `
	UserId        uint 	    `json:"userId"`
	ChannelId     uint      `json:"channelId"`
  Message       string    `json:"message" valid:"required"`
  MessageId     int       `json:"messageId"`
}

func NewHub(channelId uint) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
    channelId:  channelId,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
      for client, _ := range h.clients {
        fmt.Print("Client is %v", client)
        fmt.Print("Hub is %v", h)
        msg := Message { Nickname: client.user.UserPref.Nickname,
          Fullname: fmt.Sprintf("%s %s", client.user.UserPref.Firstname, client.user.UserPref.Lastname),
          UserId: client.user.ID,
          ChannelId: h.channelId,
          Message: fmt.Sprintf("%s has joined this channel", client.user.UserPref.Nickname),
          MessageId: int(time.Now().Unix()) }
        client.conn.WriteJSON(msg)
      }
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
      // h.broadcast <- message
      fmt.Print("\nClients ARE %v\n", len(h.clients))
			for client := range h.clients {

				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// func (h *Hub) Run() {
// 	for {
// 		select {
// 		case client := <-h.register:
// 			h.clients[client] = true
// 		case client := <-h.unregister:
// 			if _, ok := h.clients[client]; ok {
// 				delete(h.clients, client)
// 				close(client.send)
// 			}
// 		case message := <-h.broadcast:
// 			for client := range h.clients {
// 				select {
// 				case client.send <- message:
// 				default:
// 					close(client.send)
// 					delete(h.clients, client)
// 				}
// 			}
// 		}
// 	}
// }
