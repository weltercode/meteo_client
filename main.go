package main

import (
	"errors"
	"fmt"
	open_meteo_client "meteo_client/meteo_client"
	"sync"
	"time"
)

func main() {
	c, err := open_meteo_client.NewClient(time.Second*5, "Europe/Berlin")
	if err != nil {
		e := errors.New("fail to create meteo client")
		fmt.Println(e)
	}
	var results open_meteo_client.LocationResults
	results, err = c.SearchLocations("Tallinn", "EE")
	if err != nil {
		fmt.Println(err)
		return
	}
	var location open_meteo_client.LocationObj = results.Results[0]
	fmt.Println("Selected location: " + location.Name)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var forecast open_meteo_client.ForecastResponse
		forecast, err = c.GeMETNorwayForecast(location.Latitude, location.Longitude)
		if err != nil {
			fmt.Println("fail to GeMETNorwayForecast", err)
		}
		fmt.Println("GeMETNorwayForecast: ")
		fmt.Println(forecast)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		var forecast open_meteo_client.ForecastResponse
		forecast, err = c.GetTemperature(location.Latitude, location.Longitude)
		if err != nil {
			fmt.Println("fail to GetTemperature", err)
		}
		fmt.Println("GetTemperature: ")
		fmt.Println(forecast)
	}()
	wg.Wait()
}
