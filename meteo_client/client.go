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

const base_url = "https://api.open-meteo.com/v1/forecast"

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
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
	}, nil
}

func (c Client) GetTemperature(lat float32, long float32) (ForecastResponse, error) {

	var request_url string
	var forecast ForecastResponse

	request_url = fmt.Sprintf(
		"%s?latitude=%f&longitude=%f&hourly=temperature_2m&forecast_days=1&timeformat=unixtime", base_url, lat, long)
	resp, err := c.client.Get(request_url)
	if err != nil {
		return forecast, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return forecast, err
	}
	//fmt.Println(string(body))
	var r ForecastResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return forecast, err
	}
	return r, nil
}
