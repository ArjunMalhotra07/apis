package chi

import (
	"encoding/json"
	"net/http"

	"github.com/ArjunMalhotra07/apis.git/structs"
)

func BaseRoute(w http.ResponseWriter, r *http.Request) {
	response := structs.Response{Message: "pong🤣"}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the response as JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}
