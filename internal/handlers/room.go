package handlers

import (
	"fmt"
	w "meet-harbor/pkg/webrtc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RooomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/&s", guuid.New().String()))
}

func Room(c *fiber.Ctx) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return
	}
}

func Roomwebsocket(c *websocket.Conn) {
	uuid := c.Param("uuid")
	if uuid == "" {
		return
	}

	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
}

func createOrGetRoom(uuid string) (string, string, Room) {

}

func RoomViewerWebsocket(c *websocket.Conn) {

}

func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}

type websocketMessage struct {
	Event string `json:"event`
	Data  string `json:"data"`
}
