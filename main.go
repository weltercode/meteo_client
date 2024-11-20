package main

import (
	"errors"
	"fmt"
	open_meteo_client "meteo_client/meteo_client"
	"time"
)

func main() {
	var latitude, longitude float32 = 59.4165, 24.7994 // Tallinn airport
	c, err := open_meteo_client.NewClient(time.Second * 5)
	if err != nil {
		e := errors.New("fail to create meteo client")
		fmt.Println(e)
	}
	var data open_meteo_client.ForecastResponse
	data, err = c.GetTemperature(latitude, longitude)
	if err != nil {
		fmt.Println("fail to GetTemperature from meteo", err)
	}
	fmt.Println(data)
}
