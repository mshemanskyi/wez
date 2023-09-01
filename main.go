package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	var opts opts
	_, err := flags.Parse(&opts)

	if err != nil {
		fmt.Errorf("error: %s", err)
	}

	requestURL := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?q=Chernivtsi&days=1&lang=ua&key=%s", opts.ApiKey)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

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

	fmt.Printf("%s, %s\n%s \n", weather.Location.Name, weather.Location.Country, weather.Location.Localtime)
	fmt.Printf("%s %.1f°C \n\n", weather.Current.Condition.Text, weather.Current.TempC)

	for _, day := range weather.Forecast.ForecastDay {
		//fmt.Printf("%d. %s\n", i+1, day.Date)
		//fmt.Printf("   %s\n", day.Day.Condition.Text)
		//fmt.Printf("   %.1f°C\n", day.Day.AvgTempC)

		for _, hour := range day.Hour {
			date := time.Unix(hour.TimeEpoch, 0)
			if date.After(time.Now()) {
				fmt.Printf(
					"%s %.1f°C %s\n",
					date.Format("15:04"),
					hour.TempC,
					GetWeatherEmoji(hour.Condition.Text),
				)
			}
		}
		fmt.Printf("\n")
	}
}
