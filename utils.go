package logger

import (
	"encoding/json"
	"runtime"
)

// getCallerFunctionName récupère le nom de la fonction appelante
func getCallerFunctionName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// IsValidJSON checks if a string is a valid JSON
func IsValidJSON(str string) (bool, map[string]interface{}) {
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(str), &jsonData)
	if err != nil {
		return false, nil
	}
	return true, jsonData
}
