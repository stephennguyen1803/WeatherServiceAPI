package provider

import (
	"encoding/json"
	"net/http"
)

type ApiXuWeatherMap struct {
	Url string
	APIKey string
}

type ApiXuWeatherMapTemperature struct {
	Current struct {
		TemC float64 `json:"temp_c"`
	} `json:"current"`
}

func (a ApiXuWeatherMap) GetCelsiusTemperature(city string) (float64, error) {
	res, err := http.Get(a.Url + "?key=" + a.APIKey + "&q=Ho%20Chi%20Minh")

	defer res.Body.Close()

	if err != nil {
		return 0, err
	}

	weatherTemperature := ApiXuWeatherMapTemperature{}

	err = json.NewDecoder(res.Body).Decode(&weatherTemperature)

	if err != nil {
		return 0, err
	}

	return weatherTemperature.Current.TemC, nil

}
