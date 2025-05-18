package model

type SpotType string

const (
	Bicycle    SpotType = "B"
	Motorcycle SpotType = "M"
	Automobile SpotType = "A"
	Inactive   SpotType = "X"
)

type ParkingSpot struct {
	Floor         int
	Row           int
	Col           int
	SpotType      SpotType
	Active        bool
	Occupied      bool
	VehicleNumber string
}

type ParkRequest struct {
	Type          string `json:"type" example:"A"`
	VehicleNumber string `json:"vehicleNumber" example:"B1234XYZ"`
}

type UnparkRequest struct {
	SpotID        string `json:"spotId" example:"0-0-2"`
	VehicleNumber string `json:"vehicleNumber" example:"B1234XYZ"`
}
