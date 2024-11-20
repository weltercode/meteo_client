package open_meteo_client

import (
	"time"
)

type hourly struct {
	Times        []time.Time `json:"time"`
	Temperatures []float32   `json:"temperature_2m"`
}
type ForecastResponse struct {
	Latitude          float32 `json:"latitude"`
	Longitude         float32 `json:"longitude"`
	Timezone          string  `json:"timezone"`
	Timeformat        string  `json:"time"`
	Elevation         int     `json:"elevation"`
	Temeparature_sign string  `json:"temperature_2m"`
	HourlyData        hourly  `json:"hourly"`
}
