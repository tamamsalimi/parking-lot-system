package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"parking-lot/config"
	"parking-lot/handler"
	"parking-lot/service"
)

func RegisterRoutes(r *gin.Engine) {
	layout := config.GetParkingLayout()
	floors := config.GetFloors()
	rows := config.GetRows()
	cols := config.GetCols()
	rows, cols, _ = validateLayoutAndConfig(layout, floors, rows, cols)

	lot := service.NewParkingLot(floors, rows, cols, layout)
	h := handler.NewParkingHandler(lot)

	api := r.Group("/api/v1")
	{
		api.POST("/parkings", h.Park)
		api.POST("/parkings/unpark", h.Unpark)
		api.GET("/parkings/available", h.Available)
		api.GET("/parkings/search/:vehicleNumber", h.Search)
	}
}

func validateLayoutAndConfig(layout [][]string, floors, rows, cols int) (int, int, int) {
	if len(layout) == 0 || len(layout[0]) == 0 {
		log.Fatal("invalid layout: must have at least 1 row and 1 column")
	}
	if len(layout) != rows {
		log.Fatalf("invalid layout: expected %d rows, got %d", rows, len(layout))
	}
	for i, row := range layout {
		if len(row) != cols {
			log.Fatalf("invalid layout: expected %d columns in row %d, got %d", cols, i, len(row))
		}
	}

	// Constraint checks
	if floors < 1 || floors > 8 {
		log.Fatal("invalid number of floors: must be between 1 and 8")
	}
	if rows < 1 || rows > 1000 {
		log.Fatal("invalid number of rows: must be between 1 and 1000")
	}
	if cols < 1 || cols > 1000 {
		log.Fatal("invalid number of columns: must be between 1 and 1000")
	}
	if floors*rows*cols == 0 {
		log.Fatal("invalid configuration: resulting 3D layout has no capacity")
	}

	return rows, cols, cols
}
