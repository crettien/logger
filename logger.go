package logger

import (
	"fmt"
	"time"

	"github.com/crettien/logger/models"
)

// NewLogEntry crée une nouvelle entrée de log avec les détails fournis.
// Elle valide les entrées et utilise la fonction actuelle comme nom de fonction par défaut.
func NewLogEntry(level, message, source, service, tags string) (models.LogEntry, error) {
	// Valider les entrées
	if message == "" || source == "" || service == "" {
		return models.LogEntry{}, fmt.Errorf("message, source, and service must not be empty")
	}
	if level == "" {
		level = "info"
	}
	if isValid, _ := IsValidJSON(tags); tags != "" && !isValid {
		return models.LogEntry{}, fmt.Errorf("%s is not a valid JSON tag string", tags)
	}

	return models.LogEntry{
		Level:     level,
		Message:   message,
		Timestamp: time.Now(),
		Function:  getCallerFunctionName(),
		Source:    source,
		Service:   service,
		Tags:      tags,
	}, nil
}
