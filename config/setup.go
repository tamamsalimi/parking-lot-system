package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Floors int        `yaml:"floors"`
	Rows   int        `yaml:"rows"`
	Cols   int        `yaml:"cols"`
	Layout [][]string `yaml:"layout"`
}

var config Config

func init() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Println("Warning: config.yaml not found, using default config")
		setDefault()
		return
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Println("Warning: failed to parse config.yaml, using default config")
		setDefault()
		return
	}
}

func setDefault() {
	config.Floors = 2
	config.Rows = 2
	config.Cols = 3
	config.Layout = [][]string{
		{"B-1", "M-1", "A-1"},
		{"X-0", "M-1", "A-1"},
	}
}

func GetParkingLayout() [][]string {
	return config.Layout
}

func GetFloors() int {
	return config.Floors
}

func GetRows() int {
	return config.Rows
}

func GetCols() int {
	return config.Cols
}
