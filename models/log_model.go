package models

import "time"

// LogEntry représente la structure d'un log
type LogEntry struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Function  string    `json:"function"`
	Source    string    `json:"source"`
	Service   string    `json:"service"`
}
