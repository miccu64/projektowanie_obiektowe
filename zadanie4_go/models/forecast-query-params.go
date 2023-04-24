package models

type ForecastQueryParams struct {
	Latitude  float64 `query:"latitude"`
	Longitude float64 `query:"longitude"`
}
