package main

import (
	"WeatherServiceAPI/model/provider"
	"fmt"
	"time"
)

type ProviderList []provider.WeatherProvider

func main() {
	openWeather := provider.OpenWeatherMapAPI{
		"http://api.openweathermap.org/data/2.5/weather",
		"7062e5158ba18d5e0e4d6eb295a5a88f",
	}

	apixuWeather := provider.ApiXuWeatherMap{
		"http://api.apixu.com/v1/current.json",
		"ff8c4d29b9124ad9b1633538180712",
	}

	weatherBit := provider.WeatherBitMapAPI{
		"http://api.weatherbit.io/v2.0/current",
		"a5ffc6c663c246b2b8bb5fda8c02031a",
	}

	productList := ProviderList{
		openWeather,
		apixuWeather,
		weatherBit,
	}

	start := time.Now()

	//tempC, err := GetTemperatureCity("HoChiMinh", productList)

	tempC, err := GetTemperatureCity2("HoChiMinh", productList)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("nhiet do o thanh pho HCM la", tempC)

	elapsed := time.Since(start)
	fmt.Println("thoi gian xu ly la", elapsed)
}

func GetTemperatureCity(city string, providerList ProviderList) (float64, error) {

	result := float64(0)
	for _, provider := range providerList {
		tempC, err := provider.GetCelsiusTemperature(city)

		if err != nil {
			return 0, err
		}
		result = result + tempC
	}

	return result / 3, nil
}


func GetTemperatureCity2(city string, providerList ProviderList) (float64, error) {

	chanTemp := make(chan float64)
	chanErr := make(chan error)


	for _, p := range providerList {
		// Run routine
		go func(w provider.WeatherProvider) {
			temp, err := w.GetCelsiusTemperature(city)
			if err != nil {
				chanErr <- err
				return
			}
			// Đẩy dữ liệu nhiệt độ vào channel
			chanTemp <- temp
		}(p)
	}

	total := 0.0
	k := 0
	// Lấy dữ liệu nhiệt độ từ các channel (nếu có)
	for i := 0; i < len(providerList); i++ {
		select {
		case temp := <-chanTemp:
			if temp > 0 {
				total += temp
				k++
			}

		case err := <-chanErr:
			panic(err)
		}

	}
	// Sau đó tính trung bình nhiệt độ và trả kết quả
	return total / float64(k), nil
}