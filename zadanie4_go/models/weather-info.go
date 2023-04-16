package models

import (
	"gorm.io/gorm"
)

type WeatherInfo struct {
	gorm.Model
	Latitude    float32
	Longitude   float32
	Temperature float32
	Windspeed   float32
	Time        string
}
