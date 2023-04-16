package weather_proxy

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
	"sort"
	"zadanie4_go/models"
	outer_models "zadanie4_go/models/outer-models"
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
	if r.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		// https://stackoverflow.com/questions/29513472/golang-compress-flate-module-cant-decompress-valid-deflate-compressed-http-b
		bodyBytes, err = io.ReadAll(flate.NewReader(bytes.NewReader(bodyBytes[2:])))
		apiWeather := new(outer_models.ApiWeather)
		err = json.Unmarshal(bodyBytes, &apiWeather)
		if err != nil {
			return err
		}

		weather := models.WeatherInfo{
			Latitude:    apiWeather.Latitude,
			Longitude:   apiWeather.Longitude,
			Temperature: apiWeather.Current_Weather.Temperature,
			Windspeed:   apiWeather.Current_Weather.Windspeed,
			Time:        apiWeather.Current_Weather.Time,
		}
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		db.Create(&weather)

		var weathers []models.WeatherInfo
		db.Find(&weathers)
		sort.Slice(weathers, func(i, j int) bool {
			return weathers[i].ID > weathers[j].ID
		})

		// Serialize the JSON data
		jsonData, err := json.Marshal(weathers)
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
	}

	return nil
}
