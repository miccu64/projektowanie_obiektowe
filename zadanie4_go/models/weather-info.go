package models

import (
	"fmt"
	"strconv"
)

type WeatherInfo struct {
	// lat and lon are strangely rounded, so it's impossible to use them as composite key
	ID          string
	Latitude    float64
	Longitude   float64
	Temperature float64
	Windspeed   float64
	Time        string
}

func GenerateWeatherInfoID(lat float64, lon float64) string {
	return fmt.Sprintf("%s|%s", strconv.FormatFloat(lat, 'f', -1, 64), strconv.FormatFloat(lon, 'f', -1, 64))
}

func GenerateWeatherInfoIDFromString(lat string, lon string) string {
	return fmt.Sprintf("%s|%s", lat, lon)
}
