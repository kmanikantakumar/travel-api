package hotels

var hotels = []Hotel{

	{ID: "H1", Name: "Marriott City Centre", Stars: 4, PricePerNight: 189,

		Breakfast: true, Co2perNight: 12, DistanceinKm: 0.4,

		Amenities: []string{"WiFi", "Gym", "Pool"}},

	{ID: "H2", Name: "ibis Styles Central", Stars: 3, PricePerNight: 89,

		Breakfast: false, Co2perNight: 8, DistanceinKm: 1.1,

		Amenities: []string{"WiFi", "Bar"}},

	{ID: "H3", Name: "Hilton Garden Inn", Stars: 4, PricePerNight: 215,

		Breakfast: true, Co2perNight: 14, DistanceinKm: 0.8,

		Amenities: []string{"WiFi", "Gym", "Restaurant"}},
}

func SearchHotels(req HotelSearchRequest) HotelSearchResponse {

	var filteredHotels []Hotel

	for _, h := range hotels {
		if req.MinStars > 0 && h.Stars < req.MinStars {
			continue
		}

		if req.MaxPrice > 0 && h.PricePerNight > req.MaxPrice {
			continue
		}
		filteredHotels = append(filteredHotels, h)
	}

	pageSize := 10
	start := (req.Page - 1) * pageSize
	end := start + pageSize

	if end > len(filteredHotels) {
		end = len(filteredHotels)
	}

	return HotelSearchResponse{
		Hotels: filteredHotels[start:end],
		Total:  len(filteredHotels),
		Page:   req.Page,
	}
}
