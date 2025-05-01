package logger

import (
	"encoding/json"
	"regexp"
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

// IsValidKeyValuePairString checks if a string matches the format "key1:value1,key2:value2"
func IsValidKeyValuePairString(str string) bool {
	// Regex to match the format "key:value,key:value"
	pattern := `^([a-zA-Z]+:[a-zA-Z0-9.]+)(,[a-zA-Z]+:[a-zA-Z0-9.]+)*$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
