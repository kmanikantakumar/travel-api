package intent

type IntentRequest struct {
	Text string `json:"text"`
}

type TripIntent struct {
	Destination string `json:"destination"`
	DateFrom    string `json:"dateFrom"`
	DateTo      string `json:"dateTo"`
	Purpose     string `json:"purpose"`
}
