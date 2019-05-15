package api

import (
	"fmt"
	"net/http"
	"zhajinhua/server/room"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func goldFlower(c *gin.Context) {

	newConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("[ERROR]:", err)
		return
	}

	room.FGFConnChan <- newConn
}
