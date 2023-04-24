package main

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"net/url"
	"zadanie4_go/database"
	weather_proxy "zadanie4_go/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Init()

	e := echo.New()
	group := e.Group("/forecast")
	//group.GET("", controllers.GetSavedWeatherInfos)

	weatherUrl, err := url.Parse("https://api.open-meteo.com/v1")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: weatherUrl,
		},
	}

	group.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Transport:      weather_proxy.MyRoundTripper{R: http.DefaultTransport},
		ModifyResponse: weather_proxy.ModifyWeatherResponse,
		Balancer:       middleware.NewRoundRobinBalancer(targets),
		Skipper:        weather_proxy.ProxySkipper}))

	e.Logger.Fatal(e.Start(":1323"))
}
