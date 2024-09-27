package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Weather struct {
	Data []struct {
		Location struct {
			LocationName string `json:"location_name"`
		} `json:"location"`

		Date              string  `json:"date"`
		MorningForecast   string  `json:"morning_forecast"`
		AfternoonForecast string  `json:"afternoon_forecast"`
		NightForecast     string  `json:"night_forecast"`
		SummaryForecast   string  `json:"summary_forecast"`
		SummaryWhen       string  `json:"summary_when"`
		MinTemp           float64 `json:"min_temp"`
		MaxTemp           float64 `json:"max_temp"`
	} `json:"data"`
}

func main() {

	place := "Lahad Datu"

	if len(os.Args) >= 2 {
		place = os.Args[1]
	}

	today := time.Now().Format(time.DateOnly)

	req, err := http.NewRequest("GET", "https://api.data.gov.my/weather/forecast?contains="+place+"@location__location_name&filter="+today+"@date&limit=1&meta=true", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("Weather api unreachable")
	}

	var weather Weather

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		panic(err)
	}

	for _, data := range weather.Data {
		fmt.Printf("%s - %s:\n", data.Location.LocationName, data.Date)
		fmt.Printf("Pagi: %s\n", data.MorningForecast)
		fmt.Printf("Tengah Hari: %s\n", data.AfternoonForecast)
		fmt.Printf("Malam: %s\n", data.NightForecast)
		fmt.Printf("Suhu: %.2f°C - %.2f°C\n", data.MinTemp, data.MaxTemp)
		fmt.Printf("Overall: %s pada waktu %s\n", data.SummaryForecast, strings.ToLower(data.SummaryWhen))
	}

}
