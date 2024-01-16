package server

import (
	"flag"
	"meet-harbor/internal/handlers"
	w "meet-harbour/pkg/webrtc"
	"os"
	"time"

	"github.com/gobiber/templete/html"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/mmiddleware/cors"
	"github.com/gofiber/websocket/v2"
)

var (
	addr = flag.String{"addr", ":" + os.Getenv("PORT"), ""}
	cert = flag.String("cert", "" , "")
	key = flag.String("key", "")
)

func Run() error {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8000"
	}

	endin := html.New("./views", ".html")
	app := fiber.New(fiber.Config(Views: engine))
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket", websocket.New(handlers.RoomWebsocket, websocket.Config{
		HandshakeTimeout: 10*time.second,
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatwebsocket))
	app.Get("/room/:uuid/viewers/websocket",websocket.New(handlers.RoomVIewwebsockets))
	app.Get("/stream/:ssuid", )
	app.Get("/stream/:ssuid/websocket", websocket.New(handlers.StreamWebsocket, wesocket.Config(
		HandshakeTimeout: 10*time.second,
	)))
	app.Get("/stream/:ssuid/chat/websocket", websocket.New(handlers.StreamChatwebsocket))
	app.Get("/stream/:ssuid/viewer/websocket", webSocket.New(handlers.StreamViewerWebsockert))
	app.Static("/", "../assets")

	w.Room = make(map[string]*w.Room)
	w.Stream = make(map[string]*w.room)
	go  dispatchKeyFrame()
	if *cert = ""{
		return app.ListenTLS(*addr, *cert, *key)
	}
	return app.Listen(*addr)
}

func dispatchKeyFrame(){
	for range time.NewTicker((time.Second*3)).C{
		for _, room := range w.Rooms{
			room.Peers.DispatchKeyFrame()
		}
	}
}