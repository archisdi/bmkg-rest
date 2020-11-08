package repositories

import (
	"bmkg/models"
	"bmkg/utils"
	"encoding/xml"
)

// WeatherRepositoryAPI ...
type WeatherRepositoryAPI interface {
	GetLocationWeatherForecast(location string) (models.BaseWeather, error)
}

// WeatherRepository ...
type WeatherRepository struct{}

// GetWeatherForecast ...
func (*WeatherRepository) GetWeatherForecast(location string) (models.BaseWeather, error) {
	xmlBytes, err := utils.GetXMLFromURL("https://data.bmkg.go.id/datamkg/MEWS/DigitalForecast/DigitalForecast-" + location + ".xml")

	var weather models.BaseWeather
	if err = xml.Unmarshal(xmlBytes, &weather); err != nil {
		return weather, err
	}

	return weather, nil
}
