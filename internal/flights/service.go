package flights

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var allFlights = []Flight{

	{ID: "F1", Airline: "Lufthansa", FlightCode: "LH204",

		Origin: "LHR", Destination: "BER",

		Departure: "07:30", Arrival: "10:15", StopType: "Direct",

		Price: 247, Co2: 142, DurationMin: 165, SeatsLeft: 14},

	{ID: "F2", Airline: "British Airways", FlightCode: "BA782",

		Origin: "LHR", Destination: "BER",

		Departure: "09:45", Arrival: "13:00", StopType: "1 Stop",

		Price: 189, Co2: 178, DurationMin: 195, SeatsLeft: 3},

	{ID: "F3", Airline: "easyJet", FlightCode: "EZY443",

		Origin: "LHR", Destination: "BER",

		Departure: "14:00", Arrival: "17:15", StopType: "Direct",

		Price: 124, Co2: 138, DurationMin: 165, SeatsLeft: 22},
}

func SearchFlights(req FlightSearchRequest) FlightSearchResponse {
	var results []Flight

	for _, f := range allFlights {
		if req.DirectOnly && f.StopType != "Direct" {
			continue
		}
		if req.MaxPrice > 0 && f.Price > req.MaxPrice {
			continue
		}
		results = append(results, f)
	}
	return FlightSearchResponse{Flights: results, Total: len(results)}
}

func SearchFlightsList(req FlightSearchRequest) FlightSearchResponse {

	body, _ := json.Marshal(req)
	log.Println("Calling API")
	resp, err := http.Post(
		"http://localhost:8080/api/search-flights",
		"application/json",
		bytes.NewReader(body),
	)
	log.Println("API returned")

	if err != nil {
		return FlightSearchResponse{
			Flights: []Flight{},
			Total:   0,
		}
	}
	defer resp.Body.Close()

	//var result FlightSearchResponse
	//json.Unmarshal([]byte(allFlights), &result)

	return FlightSearchResponse{Flights: allFlights, Total: len(allFlights)}

}
