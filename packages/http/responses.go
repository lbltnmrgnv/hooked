package responses

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(code int32, message string) map[string]interface{} {
	return map[string]interface{}{"error": map[string]interface{}{"code": code, "message": message}}
}

func Response(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
