package provider

import (
	"encoding/json"
	"net/http"
)

type OpenWeatherMapAPI struct {
	Url string
	APIKey string
}

type OpenWeatherTemperatureMap struct {
	Current struct {
		KelvinTemp float64 `json:"temp"`
	} `json:"main"`
}

func (o OpenWeatherMapAPI) GetCelsiusTemperature(city string) (float64, error) {

	if city == "HoChiMinh" {
		city = "Thanh%20Pho%20Ho%20Chi%20Minh"
	}
	res, err := http.Get(o.Url + "?q=" + city + "&APPID=" + o.APIKey)

	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	WeatherTemperatureMap := OpenWeatherTemperatureMap{}

	err = json.NewDecoder(res.Body).Decode(&WeatherTemperatureMap)

	if err != nil {
		return 0, nil
	}

	celsiusTemperature := WeatherTemperatureMap.Current.KelvinTemp - 273.15

	return celsiusTemperature, nil

}
