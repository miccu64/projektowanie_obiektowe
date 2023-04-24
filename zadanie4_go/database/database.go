package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"zadanie4_go/models"
)

var Db *gorm.DB

func Init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = Db.AutoMigrate(&models.WeatherInfo{})
	if err != nil {
		return
	}

	weathers := []models.WeatherInfo{
		{ID: models.GenerateWeatherInfoID(11.11, 22.2), Windspeed: 1, Temperature: 2.2, Time: "2023-04-16T13:00", Longitude: 22.2, Latitude: 11.11},
		{ID: models.GenerateWeatherInfoID(11.11, 112.2), Windspeed: 3, Temperature: 21.2, Time: "2023-04-16T14:00", Longitude: 112.2, Latitude: 11.11}}
	Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&weathers)
}
