package utilSocket

import (
  "fmt"
  "time"
)

type Channel struct {

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

  channelName string

  channelType int
}

type ChannelMessage struct {
  Name        string    `json:"name"`// Name of channel
  Email       string    `json:"email"`// User Email
  Message     string    `json:"message"`// Message to send to channel
  MessageId   int       `json:"messageId"`// A random ID for the message
}

func NewChannel(channelName string, channelId uint, channelType int) *Channel {
	return &Channel{
		broadcast:    make(chan []byte),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
    channelId:    channelId,
    channelName:  channelName,
    channelType:  channelType,
	}
}

func (ch Channel) WelcomeMessage(email string) ChannelMessage {
  var msg string
  switch ch.channelType {
  case 1:
       msg = fmt.Sprintf("%v has joined this channel.", email)
       break
  case 2:
      msg = fmt.Sprintf("%v, welcome to the chatBot channel. Ask your questions", email)
      break
  case 3:
      msg = fmt.Sprintf("%v, this is your personnel channel.", email)
      break
  default:
      msg = fmt.Sprintf("%v, this is an unknown channel.", email)
  }
  return ch.SendChannelMessage(email, msg)
}

func (ch Channel) GoodbyeMessage(email string) ChannelMessage {
  msg := fmt.Sprintf("%v has left this channel.", email)
  return ch.SendChannelMessage(email, msg)
}

func (ch Channel) ProcessRandomPhrases(email string) ChannelMessage {
  msg := RandomPhrases()
  return ch.SendChannelMessage(email, msg)
}

func (ch *Channel) SendChannelMessage(email string, msg string) ChannelMessage {
  return ChannelMessage {
    Name: ch.channelName,
    Email: email,
    Message: msg,
    MessageId: int(time.Now().Unix()),
  }
}

func (ch *Channel) Run() {
	for {
		select {
		case client := <-ch.register:
			ch.clients[client] = true
      for client, _ := range ch.clients {
        msg := ch.WelcomeMessage(client.email)
        client.conn.WriteJSON(msg)
      }
		case client := <-ch.unregister:
			if _, ok := ch.clients[client]; ok {
				delete(ch.clients, client)
				close(client.send)
			}
		case message := <-ch.broadcast:
      for client := range ch.clients {
        if ch.channelType == 2 {
          msg := ch.ProcessRandomPhrases(client.email)
          client.conn.WriteJSON(msg)
        }
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(ch.clients, client)
				}
			}
		}
	}
}
