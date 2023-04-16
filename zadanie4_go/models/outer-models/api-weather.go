package outer_models

type ApiWeather struct {
	Latitude        float32
	Longitude       float32
	Current_Weather ApiWeatherNested
}
