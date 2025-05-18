package util

import "fmt"

func SpotID(floor, row, col int) string {
	return fmt.Sprintf("%d-%d-%d", floor, row, col)
}
