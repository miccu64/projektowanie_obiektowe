package outer_models

type ApiWeather struct {
	Latitude        float64
	Longitude       float64
	Current_Weather ApiWeatherNested
}
