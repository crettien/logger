package logger

import (
	"fmt"
	"net/http"

	"github.com/crettien/logger/models"
	"github.com/gorilla/websocket"
)

var WsConn *websocket.Conn

// on déclare upgrader pour une utilisation future
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Init initialise la connexion WebSocket pour l'envoi des logs
func Init(proxyURL string) error {
	conn, _, err := websocket.DefaultDialer.Dial(proxyURL, nil)
	if err != nil {
		return err
	}
	WsConn = conn
	return nil
}

// SendLog envoie un log via WebSocket
func SendLog(logEntry models.LogEntry) error {
	if WsConn == nil {
		return fmt.Errorf("%s", "WebSocket connection is not established")
	}
	return WsConn.WriteJSON(logEntry)
}
