package websocket

import (
	"net/http"
	"websocket_p4/common/log"
	"websocket_p4/websocket/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Server struct {
	clients   map[*websocket.Conn]bool
	broadcast chan models.Messages
	upgrader  websocket.Upgrader
}

func NewServer() *Server {
	server := &Server{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan models.Messages),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return server
}

func (s *Server) RunSocket(c *gin.Context) {
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Infof("error websocket : ", err)
		return
	}
	s.clients[conn] = true
	defer func() {
		delete(s.clients, conn)
		conn.Close()
	}()
	go s.handleMessages()
	for {
		var message models.Messages
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Infof("error websocket : ", err)
			break
		}
		s.broadcast <- message
	}
}

func (s *Server) handleMessages() {
	for {
		message := <-s.broadcast

		for client := range s.clients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Infof("write:", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}
