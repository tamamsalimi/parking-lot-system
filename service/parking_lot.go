package service

import (
	"errors"
	"parking-lot/model"
	"parking-lot/util"
	"strconv"
	"strings"
	"sync"
)

type ParkingLot interface {
	Park(vehicleType model.SpotType, vehicleNumber string) (string, error)
	Unpark(spotId, vehicleNumber string) error
	AvailableSpots(vehicleType model.SpotType) []string
	SearchVehicle(vehicleNumber string) (string, bool)
}

type parkingLot struct {
	sync.RWMutex
	floors     int
	rows       int
	cols       int
	spots      [][][]*model.ParkingSpot
	vehicleMap map[string]string
	available  map[model.SpotType][]*model.ParkingSpot
}

func NewParkingLot(floors, rows, cols int, layout [][]string) ParkingLot {
	spots := make([][][]*model.ParkingSpot, floors)
	available := make(map[model.SpotType][]*model.ParkingSpot)
	for f := 0; f < floors; f++ {
		spots[f] = make([][]*model.ParkingSpot, rows)
		for r := 0; r < rows; r++ {
			spots[f][r] = make([]*model.ParkingSpot, cols)
			for c := 0; c < cols; c++ {
				parts := strings.Split(layout[r][c], "-")
				typeStr := parts[0]
				active := parts[1] == "1"
				spot := &model.ParkingSpot{
					Floor:    f,
					Row:      r,
					Col:      c,
					SpotType: model.SpotType(typeStr),
					Active:   active,
				}
				spots[f][r][c] = spot
				if spot.Active {
					available[spot.SpotType] = append(available[spot.SpotType], spot)
				}
			}
		}
	}
	return &parkingLot{
		floors:     floors,
		rows:       rows,
		cols:       cols,
		spots:      spots,
		vehicleMap: make(map[string]string),
		available:  available,
	}
}

func (p *parkingLot) Park(vehicleType model.SpotType, vehicleNumber string) (string, error) {
	p.Lock()
	defer p.Unlock()

	if _, exists := p.vehicleMap[vehicleNumber]; exists {
		return "", errors.New("vehicle already parked")
	}

	queue := p.available[vehicleType]
	for i, spot := range queue {
		if !spot.Occupied {
			spot.Occupied = true
			spot.VehicleNumber = vehicleNumber
			id := util.SpotID(spot.Floor, spot.Row, spot.Col)
			p.vehicleMap[vehicleNumber] = id
			// remove from available queue
			p.available[vehicleType] = append(queue[:i], queue[i+1:]...)
			return id, nil
		}
	}
	return "", errors.New("no available spot for vehicle type")
}

func (p *parkingLot) Unpark(spotId, vehicleNumber string) error {
	p.Lock()
	defer p.Unlock()

	parts := strings.Split(spotId, "-")
	f, _ := strconv.Atoi(parts[0])
	r, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])

	if f >= p.floors || r >= p.rows || c >= p.cols {
		return errors.New("invalid spot ID")
	}
	spot := p.spots[f][r][c]
	if !spot.Occupied || spot.VehicleNumber != vehicleNumber {
		return errors.New("vehicle not found at spot")
	}
	spot.Occupied = false
	spot.VehicleNumber = ""
	p.vehicleMap[vehicleNumber] = spotId
	p.available[spot.SpotType] = append(p.available[spot.SpotType], spot)
	return nil
}

func (p *parkingLot) AvailableSpots(vehicleType model.SpotType) []string {
	p.RLock()
	defer p.RUnlock()
	var result []string
	for _, spot := range p.available[vehicleType] {
		if !spot.Occupied {
			result = append(result, util.SpotID(spot.Floor, spot.Row, spot.Col))
		}
	}
	return result
}

func (p *parkingLot) SearchVehicle(vehicleNumber string) (string, bool) {
	p.RLock()
	defer p.RUnlock()
	id, ok := p.vehicleMap[vehicleNumber]
	return id, ok
}
