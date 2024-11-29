package errorhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONError(w http.ResponseWriter, code int, field string, value string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]interface{}{
		field: value,
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println("JSONError:", err)
		return
	}
	w.Write(data)
}
