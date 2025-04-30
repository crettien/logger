package logger

import (
	"fmt"
	"time"

	"github.com/crettien/logger/models"
)

// NewLogEntry crée une nouvelle entrée de log avec les détails fournis.
// Elle valide les entrées et utilise la fonction actuelle comme nom de fonction par défaut.
func NewLogEntry(level, message, source, service string) (models.LogEntry, error) {
	// Valider les entrées
	if level == "" || message == "" || source == "" || service == "" {
		return models.LogEntry{}, fmt.Errorf("level, message, source, and service must not be empty")
	}

	return models.LogEntry{
		Level:     level,
		Message:   message,
		Timestamp: time.Now(),
		Function:  getCallerFunctionName(),
		Source:    source,
		Service:   service,
	}, nil
}
