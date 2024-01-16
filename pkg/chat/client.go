package chat

import (
	"time"

	"golang.org/fasthttp/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMassageSize = 512
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) reedPump() {

}

func (c *Client) writePump() {

}

func PeerChatConn() {

}
