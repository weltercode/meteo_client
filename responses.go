package open_meteo_client

import (
	"time"
)

type hourly struct {
	times        []time.Time `json:"time"`
	temperatures []float32   `json:"temperature_2m"`
}
type ForecastResponse struct {
	latitude          float32 `json:"latitude"`
	longitude         float32 `json:"longitude"`
	timezone          string  `json:"timezone"`
	timeformat        string  `json:"time"`
	elevation         int     `json:"elevation"`
	temeparature_sign string  `json:"temperature_2m"`
	hourlyData        hourly  `json:"hourly"`
}
