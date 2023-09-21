package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origings
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("got websocket connection", slog.String("route", "/websocket"))

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("failed to upgrade protocol to websocket", slog.String("error", err.Error()))
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNoStatusReceived) {
				slog.Info("websocket connection closed")
				return
			}
			errType := fmt.Sprintf("%T, %#v", err, err)
			slog.Error("failed to read message", slog.String("error", err.Error()), slog.String("type", errType))
			return
		}

		buff := bytes.NewBuffer(nil)
		buff.WriteString("echo: ")
		buff.Write(p)

		if err := conn.WriteMessage(messageType, buff.Bytes()); err != nil {
			slog.Error("failed to write message", slog.String("error", err.Error()))
			return
		}
	}
}
