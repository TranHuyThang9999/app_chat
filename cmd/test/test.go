package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Text      string    `json:"text"`
	UserName  string    `json:"user_name"`
	Timestamp time.Time `json:"timestamp"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func main() {
	r := gin.Default()

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		clients[conn] = true

		defer func() {
			delete(clients, conn)
			conn.Close()
		}()

		for {
			var message Message
			err := conn.ReadJSON(&message)
			if err != nil {
				log.Println("read:", err)
				break
			}

			message.Timestamp = time.Now()
			broadcast <- message
		}
	})

	go handleMessages()

	fmt.Println("server running on port 8050")
	log.Fatal(http.ListenAndServe(":8050", r))
}

func handleMessages() {
	for {
		message := <-broadcast

		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Println("write:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
