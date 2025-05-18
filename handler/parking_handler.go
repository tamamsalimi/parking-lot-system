package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"parking-lot/model"
	"parking-lot/service"
)

type ParkingHandler struct {
	Lot service.ParkingLot
}

func NewParkingHandler(lot service.ParkingLot) *ParkingHandler {
	return &ParkingHandler{Lot: lot}
}

// Park @Summary Park a vehicle
// @Param request body model.ParkRequest true "Vehicle info"
// @Success 200 {object} map[string]string
// @Failure 400,409 {object} map[string]string
// @Router /api/v1/parkings [post]
func (h *ParkingHandler) Park(c *gin.Context) {
	var req model.ParkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	spotID, err := h.Lot.Park(model.SpotType(req.Type), req.VehicleNumber)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"spotId": spotID})
}

// Unpark a vehicle
// @Param request body model.UnparkRequest true "Vehicle info"
// @Success 200 {object} map[string]string
// @Failure 400,404 {object} map[string]string
// @Router /api/v1/parkings/unpark [post]
func (h *ParkingHandler) Unpark(c *gin.Context) {
	var req model.UnparkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	err := h.Lot.Unpark(req.SpotID, req.VehicleNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "unparked"})
}

// Available @Summary Get available parking spots
// @Param type query string true "Vehicle Type (B, M, A)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/v1/parkings/available [get]
func (h *ParkingHandler) Available(c *gin.Context) {
	vehicleType := strings.ToUpper(c.Query("type"))
	if vehicleType != string(model.Bicycle) && vehicleType != string(model.Motorcycle) && vehicleType != string(model.Automobile) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vehicle type"})
		return
	}
	list := h.Lot.AvailableSpots(model.SpotType(vehicleType))
	if list == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no available spot for vehicle type"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"available": list})
}

// Search @Summary Search for parked vehicle
// @Param vehicleNumber path string true "Vehicle Number"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/parkings/search/{vehicleNumber} [get]
func (h *ParkingHandler) Search(c *gin.Context) {
	vehicleNumber := c.Param("vehicleNumber")
	spotID, ok := h.Lot.SearchVehicle(vehicleNumber)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lastSpot": spotID})
}
