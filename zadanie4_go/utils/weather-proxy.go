package weather_proxy

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"zadanie4_go/database"
	"zadanie4_go/models"
	outerModels "zadanie4_go/models/outer-models"
)

type MyRoundTripper struct {
	R http.RoundTripper
}

// RoundTrip edit initial request before proxy
func (mrt MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.URL.RawQuery += "&current_weather=true"
	return mrt.R.RoundTrip(r)
}

func ModifyWeatherResponse(r *http.Response) error {
	if r.StatusCode != http.StatusOK {
		return nil
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err = io.ReadAll(flate.NewReader(bytes.NewReader(bodyBytes[2:])))
	apiWeather := new(outerModels.ApiWeather)
	err = json.Unmarshal(bodyBytes, &apiWeather)
	if err != nil {
		return err
	}

	weather := models.WeatherInfo{
		ID:          models.GenerateWeatherInfoIDFromString(r.Request.URL.Query().Get("latitude"), r.Request.URL.Query().Get("longitude")),
		Latitude:    apiWeather.Latitude,
		Longitude:   apiWeather.Longitude,
		Temperature: apiWeather.Current_Weather.Temperature,
		Windspeed:   apiWeather.Current_Weather.Windspeed,
		Time:        apiWeather.Current_Weather.Time,
	}
	database.Db.Create(&weather)

	// Serialize the JSON data
	jsonData, err := json.Marshal(weather)
	if err != nil {
		return err
	}

	// Compress the JSON data using deflate
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		return err
	}
	if _, err := writer.Write(jsonData); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	jsonData = buf.Bytes()

	// Update the response body with the compressed JSON data
	r.Body = io.NopCloser(bytes.NewReader(jsonData))
	r.ContentLength = int64(len(jsonData))

	if err != nil {
		return err
	}
	return nil
}

func ProxySkipper(c echo.Context) bool {
	queryParams := new(models.ForecastQueryParams)
	err := c.Bind(queryParams)
	if err != nil {
		return true
	}

	result := models.WeatherInfo{ID: models.GenerateWeatherInfoID(queryParams.Latitude, queryParams.Longitude)}
	err = database.Db.First(&result).Error
	if err == nil {
		err = c.JSON(http.StatusOK, result)
		if err != nil {
			return false
		}
		return true
	}
	return false
}
