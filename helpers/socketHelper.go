package helpers

// import (
//   "fmt"
//   "net/http"
//   "github.com/gorilla/websocket"
// )
//
// type Message struct {
// 	Nickname      string    `json:"nickname"`
// 	Fullname      string    `json:"fullname" `
// 	UserId        uint 	    `json:"userId"`
// 	ChannelId     uint      `json:"channelId"`
//   Message       string    `json:"message" valid:"required"`
//   MessageId     int       `json:"messageId"`
// }
//
// var socketUpgrader = websocket.Upgrader{
//   CheckOrigin: func(r *http.Request) bool {
//     return true
//   },
//   ReadBufferSize:  1024,
//   WriteBufferSize: 1024,
// }

// func HandleIncoming(w http.ResponseWriter, r *http.Request, isChatBot bool) {
//   conn, err := socketUpgrader.Upgrade(w, r, nil)
//     if err != nil {
//         fmt.Println("Failed to set websocket upgrade: %v", err)
//         return
//     }
//
//     for {
//         t, msg, err := conn.ReadMessage()
//         if err != nil {
//             break
//         }
//         conn.WriteMessage(t, msg)
//         fmt.Printf("\nreceived:%v\n", msg)
//         if isChatBot {
//           // Process Chatbot response
//           // ProcessChatbotResponse(msg)
//         }
//     }
// }
