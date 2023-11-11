package configuration

type Timeouts struct {
	Connection float64
	Response   float64
}

type Retries struct {
	Count    int
	OnErrors []int
	Delay    float64
}

type Http struct {
	Timeouts Timeouts
	Retries  Retries
}

type Config struct {
	Host             string
	Endpoint         string
	DefaultCountries string
	UserAgent        string
	Http             Http
	LogLevel         string
}

func GetConfig() *Config {
	return &Config{
		Host:             "https://randomuser.me",
		Endpoint:         "/api",
		DefaultCountries: "us,gb,au,ca,ie",
		UserAgent:        "random-user",
		Http: Http{
			Timeouts: Timeouts{
				Connection: 5,
				Response:   10,
			},
			Retries: Retries{
				Count:    3,
				OnErrors: []int{500, 502, 503},
				Delay:    0.5,
			},
		},
		LogLevel: "info",
	}
}
