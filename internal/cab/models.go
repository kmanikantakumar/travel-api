package cab

import "errors"

type CabBookingRequest struct {
	TripID string `json:"tripId"`

	PickupAddr string `json:"pickupAddress"`

	DropoffAddr string `json:"dropoffAddress"`

	PickupTime string `json:"pickupTime"` // '2026-05-22T06:00:00'

	VehicleType string `json:"vehicleType"` // 'Saloon'|'Estate'|'MPV'|'Executive'

	Passengers int `json:"passengers"`
}

type CabBookingResponse struct {
	BookingID string `json:"bookingId"`

	Status string `json:"status"`

	DriverName string `json:"driverName"`

	VehiclePlate string `json:"vehiclePlate"`

	EstPrice int `json:"estimatedPrice"`
}

func ValidateCabBookingRequest(req CabBookingRequest) error {

	if req.PickupAddr == "" {
		return errors.New("pickupAddress is required")
	}
	if req.DropoffAddr == "" {
		return errors.New("dropoffAddress is required")
	}
	if req.PickupTime == "" {
		return errors.New("pickupTime is required")
	}
	if req.VehicleType == "" {
		return errors.New("vehicleType is required")
	}
	if req.Passengers <= 0 {
		return errors.New("passengers must be atleast 1")
	}

	return nil
}
