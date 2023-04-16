package models

type ForecastQueryParams struct {
	Latitude  float32 `query:"latitude"`
	Longitude float32 `query:"longitude"`
}
