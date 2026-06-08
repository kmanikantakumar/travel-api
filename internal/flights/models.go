package flights

type FlightSearchRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	DirectOnly  bool   `json:"directOnly"`
	MaxPrice    int    `json:"maxPrice"`
}

type Flight struct {
	ID string `json:"id"`

	Airline string `json:"airline"`

	FlightCode string `json:"flightCode"`

	Origin string `json:"origin"`

	Destination string `json:"destination"`

	Departure string `json:"departure"` // '07:30'

	Arrival string `json:"arrival"` // '10:15'

	StopType string `json:"stopType"` // 'Direct' | '1 Stop'

	Price int `json:"price"`

	Co2 int `json:"co2"`

	DurationMin int `json:"durationMin"`

	SeatsLeft int `json:"seatsLeft"`
}

type FlightSearchResponse struct {
	Flights []Flight `json:"flights"`
	Total   int      `json:"total"`
}
