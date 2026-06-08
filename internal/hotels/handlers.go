package hotels

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func SearchHotelsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	city := query.Get("city")
	checkIn := query.Get("check_in")
	checkOut := query.Get("check_out")

	stars, _ := strconv.Atoi(query.Get("stars"))
	maxPrice, _ := strconv.Atoi(query.Get("max_price"))
	page, _ := strconv.Atoi(query.Get("page"))

	if page < 1 {
		page = 1
	}

	res := SearchHotels(HotelSearchRequest{
		City:     city,
		CheckIn:  checkIn,
		CheckOut: checkOut,
		MinStars: stars,
		MaxPrice: maxPrice,
		Page:     page,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
