package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("wez")

	var opts opts
	_, err := flags.Parse(&opts)

	if err != nil {
		fmt.Errorf("error: %s", err)
	}

	requestURL := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?q=Chernivtsi&lang=ua&key=%s", opts.ApiKey)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s\n", err)
		os.Exit(1)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("error unmarshalling json: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: weather: %+v\n", weather)
	fmt.Println("weather.Current.Condition: %s", weather.Current.Condition)
}
