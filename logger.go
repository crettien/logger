package logger

import (
	"time"

	"github.com/crettien/logger/models"
)

// CreateLogEntry crée une nouvelle entrée de log
func CreateLogEntry(level, message, source, service string) models.LogEntry {
	return models.LogEntry{
		Level:     level,
		Message:   message,
		Timestamp: time.Now(),
		Function:  getCallerFunctionName(),
		Source:    source,
		Service:   service,
	}
}
