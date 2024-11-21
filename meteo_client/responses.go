package open_meteo_client

import (
	"encoding/json"
	"time"
)

/*
		Example json
	  "latitude": 52.52,
	  "longitude": 13.419,
	  "elevation": 44.812,
	  "generationtime_ms": 2.2119,
	  "utc_offset_seconds": 0,
	  "timezone": "Europe/Berlin",
	  "timezone_abbreviation": "CEST",
	  "hourly": {
	    "time": ["2022-07-01T00:00", "2022-07-01T01:00", "2022-07-01T02:00", ...],
	    "temperature_2m": [13, 12.7, 12.7, 12.5, 12.5, 12.8, 13, 12.9, 13.3, ...]
	  },
	  "hourly_units": {
	    "temperature_2m": "°C"
	  }
*/

func (h *hourly) UnmarshalJSON(data []byte) error {
	type Alias hourly
	aux := &struct {
		Times []int64 `json:"time"` // Use `int64` to parse Unix timestamps
		*Alias
	}{
		Alias: (*Alias)(h),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Convert Unix timestamps to `time.Time`
	for _, t := range aux.Times {
		h.Times = append(h.Times, time.Unix(t, 0))
	}

	return nil
}

type hourly struct {
	Times        []time.Time `json:"time"`
	Temperatures []float32   `json:"temperature_2m"`
}
type ForecastResponse struct {
	Latitude          float32 `json:"latitude"`
	Longitude         float32 `json:"longitude"`
	Timezone          string  `json:"timezone"`
	Timeformat        string  `json:"time"`
	Elevation         float32 `json:"elevation"`
	Temeparature_sign string  `json:"temperature_2m"`
	HourlyData        hourly  `json:"hourly"`
}

type LocationObj struct {
	Id          int32   `json:"id"`
	Name        string  `json:"name"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	Elevation   float32 `json:"elevation"`
	FeatureCode string  `json:"feature_code"`
	CountryCode string  `json:"country_code"`
	AdminId1    int32   `json:"admin1_id"`
	Timezone    string  `json:"timezone"`
	Population  int32   `json:"population"`
	Country_id  int32   `json:"country_id"`
	Country     string  `json:"country"`
	Admin1      string  `json:"admin1"`
}
type LocationResults struct {
	Results []LocationObj `json:"results"`
}
