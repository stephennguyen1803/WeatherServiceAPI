package provider

import (
	"encoding/json"
	"net/http"
)

type WeatherBitMapAPI struct {
	Url string
	APIKey string
}

type WeatherBitMapTemperature struct {
	Current []struct {
		TempC float64 `json:"temp"`
	} `json:"data"`
}

func (w WeatherBitMapAPI) GetCelsiusTemperature(city string) (float64, error) {
	res, err := http.Get(w.Url + "?city=" + city + "&key=" + w.APIKey)

	defer res.Body.Close()

	if err != nil {
		return 0, nil
	}

	weatherTemperature := WeatherBitMapTemperature{}

	err = json.NewDecoder(res.Body).Decode(&weatherTemperature)

	if err != nil {
		return 0, err
	}

	return weatherTemperature.Current[0].TempC, nil

}
