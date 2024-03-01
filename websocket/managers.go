package websocket

import (
	"github.com/gin-gonic/gin"
)

type ManagerClient struct {
	sockets map[string]*server
}

func NewManagerClient() *ManagerClient {
	return &ManagerClient{
		sockets: make(map[string]*server),
	}
}
func (mc *ManagerClient) ServerWs(c *gin.Context) {

	room := c.Query("room")

	if _, ok := mc.sockets[room]; !ok {
		mc.sockets[room] = NewServer() // init room
	}
	mc.sockets[room].runSocket(c, room)
}
