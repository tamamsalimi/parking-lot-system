package test

import (
	"parking-lot/model"
	"parking-lot/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLot() service.ParkingLot {
	layout := [][]string{
		{"B-1", "M-1", "A-1"},
		{"X-0", "M-1", "A-1"},
	}
	return service.NewParkingLot(2, 2, 3, layout)
}

func TestParkAndUnpark(t *testing.T) {
	lot := getLot()
	spotID, err := lot.Park(model.Motorcycle, "M001")
	assert.NoError(t, err)
	assert.NotEmpty(t, spotID)

	err = lot.Unpark(spotID, "M001")
	assert.NoError(t, err)
}

func TestParkFailWhenFull(t *testing.T) {
	lot := getLot()
	_, _ = lot.Park(model.Bicycle, "B001")
	_, _ = lot.Park(model.Bicycle, "B002")
	_, err := lot.Park(model.Bicycle, "B003")
	assert.Error(t, err)
}
func TestDuplicateParkingFails(t *testing.T) {
	lot := getLot()
	_, err := lot.Park(model.Automobile, "DUP123")
	assert.NoError(t, err)
	_, err = lot.Park(model.Automobile, "DUP123")
	assert.Error(t, err)
	assert.Equal(t, "vehicle already parked", err.Error())
}

func TestUnparkInvalid(t *testing.T) {
	lot := getLot()
	err := lot.Unpark("0-0-1", "WRONG")
	assert.Error(t, err)
}

func TestUnparkFromEmptySpot(t *testing.T) {
	lot := getLot()
	err := lot.Unpark("0-0-1", "")
	assert.Error(t, err)
}

func TestAvailableSpots(t *testing.T) {
	lot := getLot()
	spots := lot.AvailableSpots(model.Automobile)
	assert.Len(t, spots, 4)
}

func TestSearchVehicle(t *testing.T) {
	lot := getLot()
	spotID, _ := lot.Park(model.Automobile, "A001")
	_ = lot.Unpark(spotID, "A001")
	last, found := lot.SearchVehicle("A001")
	assert.True(t, found)
	assert.Equal(t, spotID, last)
}

func TestSearchVehicleNotFound(t *testing.T) {
	lot := getLot()
	_, found := lot.SearchVehicle("UNKNOWN")
	assert.False(t, found)
}
