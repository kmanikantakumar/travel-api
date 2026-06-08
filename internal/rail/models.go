package rail

type RailSearchRequest struct {
	Origin string `json:"origin"`

	Destination string `json:"destination"`

	Date string `json:"date"`

	SortBy string `json:"sortBy"` // 'price' | 'duration' | 'departure'

}

type RailClass struct {
	Name string `json:"name"`

	Price int `json:"price"`

	Availability int `json:"availability"`
}

type RailService struct {
	ID string `json:"id"`

	Operator string `json:"operator"`

	TrainCode string `json:"trainCode"`

	Departure string `json:"departure"`

	Arrival string `json:"arrival"`

	DurationMin int `json:"durationMin"`

	Co2 int `json:"co2"`

	Classes []RailClass `json:"classes"`
}
