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
	request_url = fmt.Sprintf(
		"%s?latitude=%f&longitude=%f&hourly=temperature_2m", base_url, lat, long)
	resp, err := c.client.Get(request_url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r ForecastResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return r, nil
}
