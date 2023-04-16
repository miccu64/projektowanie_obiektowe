package main

import (
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"zadanie4_go/controllers"
	"zadanie4_go/models"
	weather_proxy "zadanie4_go/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.WeatherInfo{})
	if err != nil {
		return
	}

	weathers := []models.WeatherInfo{{Windspeed: 1, Temperature: 2.2, Time: "2023-04-16T13:00", Longitude: 22.2, Latitude: 11.11},
		{Windspeed: 3, Temperature: 21.2, Time: "2023-04-16T14:00", Longitude: 112.2, Latitude: 11.11}}
	db.Create(&weathers)

	e := echo.New()
	group := e.Group("/forecast")
	group.GET("", controllers.GetSavedWeatherInfos)

	weatherUrl, err := url.Parse("https://api.open-meteo.com/v1")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: weatherUrl,
		},
	}

	group.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{Transport: weather_proxy.MyRoundTripper{R: http.DefaultTransport}, ModifyResponse: weather_proxy.ModifyWeatherResponse, Balancer: middleware.NewRoundRobinBalancer(targets)}))

	e.Logger.Fatal(e.Start(":1323"))
}
