package hotels

type HotelSearchRequest struct {
	City     string `json:"city"`
	CheckIn  string `json:"checkIn"`  // '2026-05-22'
	CheckOut string `json:"checkOut"` // '2026-05-25'
	MinStars int    `json:"minStars"` // 1-5
	MaxPrice int    `json:"maxPrice"` // in inr
	Page     int    `json:"page"`     // for pagination
}

type Hotel struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Stars         int      `json:"stars"`
	PricePerNight int      `json:"priceperNight"` // per night in inr
	Breakfast     bool     `json:"breakfast"`
	Co2perNight   int      `json:"co2perNight"`  // in kg
	DistanceinKm  float64  `json:"distanceInKm"` // from city center
	Amenities     []string `json:"amenities"`
}

type HotelSearchResponse struct {
	Hotels []Hotel `json:"hotels"`
	Total  int     `json:"total"`
	Page   int     `json:"page"`
}
