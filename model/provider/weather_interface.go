package provider

type WeatherProvider interface {
	GetCelsiusTemperature (string)(float64, error)
}


