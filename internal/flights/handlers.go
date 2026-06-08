package flights

import (
	"encoding/json"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	// req := FlightSearchRequest{
	// 	Origin:      "LHR",
	// 	Destination: "BER",
	// 	Date:        "2026-05-22",
	// 	DirectOnly:  true,
	// 	MaxPrice:    300,
	// }
	var req FlightSearchRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	res := SearchFlights(req)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}
