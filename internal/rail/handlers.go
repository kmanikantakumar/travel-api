package rail

import (
	"encoding/json"
	"net/http"
)

func SearchRailsHandler(w http.ResponseWriter, r *http.Request) {

	var req RailSearchRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	res := SearchRail(req)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)

}
