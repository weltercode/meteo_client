package open_meteo_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
API docs for Weather Forecast
https://open-meteo.com/en/docs
*/
const weather_base_url = "https://api.open-meteo.com/v1/forecast"
const location_base_url = "https://geocoding-api.open-meteo.com/v1/search"

type Client struct {
	client   *http.Client
	timezone string
}

func NewClient(timeout time.Duration, timezone string) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
		timezone: timezone,
	}, nil
}

func (c Client) SearchLocations(location_name string, country_code string) (LocationResults, error) {
	var results LocationResults
	request_url := fmt.Sprintf("%s?name=%s", location_base_url, location_name)
	err := c.parseResponse(request_url, &results)
	var new_result LocationResults
	for _, loc := range results.Results {
		if loc.CountryCode == country_code {
			new_result.Results = append(new_result.Results, loc)
		}
	}
	return new_result, err
}

func (c Client) GetTemperature(lat float32, long float32) (ForecastResponse, error) {
	return c.getForecast(lat, long, 1, "")
}

func (c Client) GeMETNorwayForecast(lat float32, long float32) (ForecastResponse, error) {
	return c.getForecast(lat, long, 1, "gfs_seamless")
}

func (c Client) getForecast(lat float32, long float32, forecast_days int32, models string) (ForecastResponse, error) {

	var forecast ForecastResponse
	request_url := fmt.Sprintf("%s?latitude=%f&longitude=%f&hourly=temperature_2m&timezone=%s&forecast_days=%d&timeformat=unixtime",
		weather_base_url, lat, long, c.timezone, forecast_days)
	if models != "" {
		request_url += fmt.Sprintf("&models=%s", models)
	}
	err := c.parseResponse(request_url, &forecast)
	return forecast, err
}

func (c Client) parseResponse(request_url string, responseObj interface{}) error {
	resp, err := c.client.Get(request_url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, responseObj); err != nil {
		return err
	}
	return nil
}
