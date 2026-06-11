package trips

type Trip struct {
	ID          string `json:"id"`
	TravellerID string `json:"travellerId"`
	Destination string `json:"destination"`
	DateFrom    string `json:"dateFrom"`
	DateTo      string `json:"dateTo"`
	Purpose     string `json:"purpose"`
	Status      string `json:"status"`
	TotalCost   int    `json:"totalCost"`
	Notes       string `json:"notes"`
}

type TripSegment struct {
	ID          int64  `json:"id"`
	TripID      string `json:"tripId"`
	SegmentType string `json:"segmentType"`
	RefID       string `json:"refId"`
	Details     string `json:"details"`
	Cost        int    `json:"cost"`
}

type CreateTripRequest struct {
	TravellerID string        `json:"travellerId"`
	Destination string        `json:"destination"`
	DateFrom    string        `json:"dateFrom"`
	DateTo      string        `json:"dateTo"`
	Purpose     string        `json:"purpose"`
	TotalCost   int           `json:"totalCost"`
	Notes       string        `json:"notes"`
	Segments    []TripSegment `json:"segments"`
}
