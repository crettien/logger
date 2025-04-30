package logger

import (
	"fmt"
	"strings"

	"github.com/crettien/logger/models"
	"github.com/gorilla/websocket"
)

var wsConn *websocket.Conn

// InitializeWebSocketConnection initialise la connexion WebSocket pour l'envoi des logs
func InitializeWebSocketConnection(proxyURL string) error {
	if !strings.HasPrefix(proxyURL, "ws://") && !strings.HasPrefix(proxyURL, "wss://") {
		return fmt.Errorf("%s is not a valid websocket URL", proxyURL)
	}
	conn, _, err := websocket.DefaultDialer.Dial(proxyURL, nil)
	if err != nil {
		return err
	}
	wsConn = conn
	return nil
}

// SendLogOverWebSocket envoie un log via WebSocket
func SendLogOverWebSocket(logEntry models.LogEntry) error {
	if wsConn == nil {
		return fmt.Errorf("WebSocket connection is not established")
	}
	return wsConn.WriteJSON(logEntry)
}

// CloseWebSocketConnection ferme la connexion WebSocket si elle est ouverte
func CloseWebSocketConnection() error {
	if wsConn != nil {
		err := wsConn.Close()
		if err != nil {
			return err
		}
		wsConn = nil
	}
	return nil
}

// IsWebSocketConnected vérifie si la connexion WebSocket est établie
func IsWebSocketConnected() bool {
	return wsConn != nil
}
