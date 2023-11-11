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
	Api       string
	UserAgent string
	Http      Http
}
