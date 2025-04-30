//go:build test
// +build test

package logger

import "github.com/gorilla/websocket"

// GetWsConn exposes wsConn for testing only
func GetWsConn() *websocket.Conn {
	return wsConn
}
