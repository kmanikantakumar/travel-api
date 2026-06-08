package cab

import (
	"fmt"
	"math/rand"
)

var basePrices = map[string]int{

	"Saloon": 35, "Estate": 42, "MPV": 55, "Executive": 80,
}

func BookCab(req CabBookingRequest) CabBookingResponse {

	bookingID := fmt.Sprintf("CAB-%04X", rand.Intn(65536))
	price := basePrices[req.VehicleType]
	if price == 0 {
		price = 40 // default price for unknown vehicle types
	}

	return CabBookingResponse{

		BookingID: bookingID, Status: "Confirmed", DriverName: "John Doe", VehiclePlate: "XYZ 1234", EstPrice: price,
	}
}
