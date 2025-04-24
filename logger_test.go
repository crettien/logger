package logger_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/crettien/logger"
	"github.com/crettien/logger/models"
	"github.com/gorilla/websocket"
)

func TestInitWebSocket(t *testing.T) {
	// Créer un serveur WebSocket pour le test
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		defer conn.Close()
	}))
	defer server.Close()

	// Remplacer "http" par "ws" pour une URL WebSocket correcte
	wsURL := "ws" + server.URL[4:] + "/ws"

	// Initialiser la connexion WebSocket
	err := logger.Init(wsURL)
	if err != nil {
		t.Fatalf("Failed to initialize WebSocket connection: %v", err)
	}

	// Vérifier que la connexion est établie
	if logger.WsConn == nil {
		t.Fatal("WebSocket connection is not established")
	}
}

func TestCreateLogEntry(t *testing.T) {
	logEntry := logger.CreateLogEntry("info", "Test log message", "test-source", "test-service")

	if logEntry.Level != "info" {
		t.Errorf("Expected log level to be 'info', got %s", logEntry.Level)
	}
	if logEntry.Message != "Test log message" {
		t.Errorf("Expected log message to be 'Test log message', got %s", logEntry.Message)
	}
	if logEntry.Source != "test-source" {
		t.Errorf("Expected log source to be 'test-source', got %s", logEntry.Source)
	}
	if logEntry.Service != "test-service" {
		t.Errorf("Expected log service to be 'test-service', got %s", logEntry.Service)
	}
	if logEntry.Function == "" {
		t.Error("Expected log function to be not empty")
	}
	if time.Since(logEntry.Timestamp) > 5*time.Second {
		t.Error("Expected log timestamp to be recent")
	}
}

func TestSendLog(t *testing.T) {
	// Créer un serveur WebSocket pour le test
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		defer conn.Close()

		// Lire le message envoyé par le client
		_, message, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("Failed to read message: %v", err)
		}

		// Analyser le message JSON
		var logEntry models.LogEntry
		err = json.Unmarshal(message, &logEntry)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON message: %v", err)
		}

		// Vérifier le contenu du message
		if logEntry.Message != "Test log message" {
			t.Errorf("Expected message to be 'Test log message', got %s", logEntry.Message)
		}
	}))
	defer server.Close()

	// Remplacer "http" par "ws" pour une URL WebSocket correcte
	wsURL := "ws" + server.URL[4:]

	// Initialiser la connexion WebSocket
	err := logger.Init(wsURL)
	if err != nil {
		t.Fatalf("Failed to initialize WebSocket connection: %v", err)
	}

	// Créer et envoyer un log
	logEntry := logger.CreateLogEntry("info", "Test log message", "test-source", "test-service")
	err = logger.SendLog(logEntry)
	if err != nil {
		t.Fatalf("Failed to send log: %v", err)
	}
}
