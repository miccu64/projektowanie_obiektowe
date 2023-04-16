package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"zadanie4_go/models"
)

func GetSavedWeatherInfos(c echo.Context) (err error) {
	queryParams := new(models.ForecastQueryParams)
	if err = c.Bind(queryParams); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	var weathers []models.WeatherInfo
	db.Find(&weathers)

	return c.JSON(http.StatusOK, weathers)
}
