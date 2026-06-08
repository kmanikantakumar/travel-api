package rail

import (
	"sort"
)

var routeData = map[string][]RailService{
	"LHR-AMS": {

		{ID: "R1", Operator: "Eurostar", TrainCode: "ES9128",

			Departure: "06:00", Arrival: "11:45", DurationMin: 345, Co2: 18,

			Classes: []RailClass{

				{Name: "Standard", Price: 89, Availability: 42},

				{Name: "First Class", Price: 189, Availability: 8},
			}},

		{ID: "R2", Operator: "Thalys", TrainCode: "TH114",

			Departure: "09:00", Arrival: "14:30", DurationMin: 330, Co2: 16,

			Classes: []RailClass{

				{Name: "Standard", Price: 112, Availability: 18},

				{Name: "First Class", Price: 220, Availability: 4},
			}},
	},
}

func SearchRail(req RailSearchRequest) []RailService {
	key := req.Origin + "-" + req.Destination
	services, exists := routeData[key]

	if services == nil || !exists {
		return []RailService{}
	}

	sort.Slice(services, func(i, j int) bool {
		switch req.SortBy {
		case "duration":
			return services[i].DurationMin < services[j].DurationMin
		case "departure":
			return services[i].Departure < services[j].Departure
		default:
			return services[i].Classes[0].Price < services[j].Classes[0].Price

		}
	})
	return services
}
