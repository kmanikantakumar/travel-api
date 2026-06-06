package intent

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var request IntentRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	result := Parse(request.Text)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}
