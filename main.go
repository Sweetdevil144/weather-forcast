package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	} `json:"location"`
	Current struct {
		TempC      float64 `json:"temp_c"`
		WindMph    float64 `json:"wind_mph"`
		WindDegree float64 `json:"wind_degree"`
		Pressure   float64 `json:"pressure"`
		PrecipMm   float64 `json:"precip_mm"`
		Cloud      float64 `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		HeatindexC float64 `json:"heatindex_c"`
		Condition  struct {
			Text string `json:"text"`
		} `json:"condition"`
		DewPoint float64 `json:"dawpoint_c"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Date string `json:"date"`
			Day  struct {
				MaxtempC      float64 `json:"maxtemp_c"`
				MintempC      float64 `json:"mintemp_c"`
				AvgtempC      float64 `json:"avgtemp_c"`
				MaxwindMph    float64 `json:"maxwind_mph"`
				TotalprecipMm float64 `json:"totalprecip_mm"`
				TotalsnowCm   float64 `json:"totalsnow_cm"`
				AvgHumidity   float64 `json:"avghumidity"`
				Condition     struct {
					Text string `json:"text"`
				} `json:"condition"`
				UV float64 `json:"uv"`
			} `json:"day"`
			Hour []struct {
				Time      string `json:"time"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				FeelslikeC float64 `json:"feelslike_c"`
				Humidity   float64 `json:"humidity"`
				Cloud      float64 `json:"cloud"`
				WillItRain int32   `json:"will_it_rain"`
				WindMph    float64 `json:"wind_mph"`
				WillItSnow float64 `json:"will_it_snow"`
				UV         float64 `json:"uv"`
			} `json:"hour"`
			Astro struct {
				Sunrise  string `json:"sunrise"`
				Sunset   string `json:"sunset"`
				Moonrise string `json:"moonrise"`
				Moonset  string `json:"moonset"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	var weather Weather
	var input string

	cyan := color.New(color.FgCyan).PrintlnFunc()
	red := color.New(color.FgRed).PrintlnFunc()
	green := color.New(color.FgGreen).PrintlnFunc()
	yellow := color.New(color.FgYellow).PrintlnFunc()
	blue := color.New(color.FgBlue).PrintlnFunc()
	purple := color.New(color.FgMagenta).PrintlnFunc()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}

	if len(os.Args) >= 2 {
		input = os.Args[1]
	} else {
		purple("Enter your city name")
		fmt.Scanln(&input)
		if err != nil {
			log.Fatal("An Error occured :", err)
		}
		if input == "" {
			// Empty Input
			input = "London" // Defaulting to London
			red("No City name Provided. Choose Any from the below list\n")
			blue("1.Varanasi\n2.Paris\n3.New York\n4.Moscow")
			var option string
			fmt.Scanln(&option)
			if option == "" {
				yellow("No Options chosen. You lazy Bastard :/ :/ ........ \nDefaulting to London ........")
				input = "London"
			} else {
				if option == "1" {
					input = "Varanasi"
				} else if option == "2" {
					input = "Paris"
				} else if option == "3" {
					input = "New York"
				} else if option == "4" {
					input = "Moscow"
				} else {
					red("My Champagne Lover!! Please choose a correct option next time. Will you? :/ \n Defaulting to Varanasi !! :/ .....\n")
					time.Sleep(2 * time.Second)
					input = "Varanasi"
				}
			}
		}
	}

	query := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", apiKey, input)
	res, err := http.Get(query)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("ResPonse Code != 200")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &weather)

	heading := fmt.Sprintf("\n------------------------------More Information on %s------------------------------\n", strings.ToUpper(input))
	red(heading)
	cyan("Location: ", weather.Location.Name)
	cyan("Region: ", weather.Location.Region)
	cyan("Country: ", weather.Location.Country)
	cyan("Latitude: ", weather.Location.Lat)
	cyan("Longitude: ", weather.Location.Lon)
	cyan("Temperature: ", weather.Current.TempC)
	cyan("Wind Speed: ", weather.Current.WindMph)
	cyan("Wind Degree: ", weather.Current.WindDegree)
	cyan("Pressure: ", weather.Current.Pressure)
	cyan("Precipitation: ", weather.Current.PrecipMm)
	cyan("Cloud: ", weather.Current.Cloud)
	cyan("Feels Like: ", weather.Current.FeelslikeC)
	cyan("Heat Index: ", weather.Current.HeatindexC)
	cyan("Condition: ", weather.Current.Condition.Text)
	cyan("Dew Point: ", weather.Current.DewPoint)
	red("\n------------------------------Forecast------------------------------\n")
	if len(weather.Forecast.ForecastDay) > 0 {
		yellow("Max Daily Temperature", weather.Forecast.ForecastDay[0].Day.MaxtempC)
		yellow("Min Daily Temperature ", weather.Forecast.ForecastDay[0].Day.MintempC)
		yellow("Average Daily Temperature ", weather.Forecast.ForecastDay[0].Day.AvgtempC)
		yellow("Max Wind Speed ", weather.Forecast.ForecastDay[0].Day.MaxwindMph)
		yellow("Total Precipitation ", weather.Forecast.ForecastDay[0].Day.TotalprecipMm)
		yellow("Total Snow ", weather.Forecast.ForecastDay[0].Day.TotalsnowCm)
		yellow("Average Humidity ", weather.Forecast.ForecastDay[0].Day.AvgHumidity)
		yellow("Condition ", weather.Forecast.ForecastDay[0].Day.Condition.Text)
		yellow("UV ", weather.Forecast.ForecastDay[0].Day.UV)
	}
	red("\n------------------------------Astro------------------------------\n")
	if len(weather.Forecast.ForecastDay) > 0 {
		green("Sunrise ", weather.Forecast.ForecastDay[0].Astro.Sunrise)
		green("Sunset ", weather.Forecast.ForecastDay[0].Astro.Sunset)
		green("Moonrise ", weather.Forecast.ForecastDay[0].Astro.Moonrise)
		green("Moonset ", weather.Forecast.ForecastDay[0].Astro.Moonset)
	}
	green("\n------------------------------Hourly Forecast------------------------------\n")
	for i := 0; i < 24; i++ {
		blue("Time: ", weather.Forecast.ForecastDay[0].Hour[i].Time)
		purple("Condition: ", weather.Forecast.ForecastDay[0].Hour[i].Condition.Text)
		blue("Feels Like: ", weather.Forecast.ForecastDay[0].Hour[i].FeelslikeC)
		purple("Humidity: ", weather.Forecast.ForecastDay[0].Hour[i].Humidity)
		blue("Cloud: ", weather.Forecast.ForecastDay[0].Hour[i].Cloud)
		purple("Will It Rain: ", weather.Forecast.ForecastDay[0].Hour[i].WillItRain)
		blue("Wind Speed: ", weather.Forecast.ForecastDay[0].Hour[i].WindMph)
		purple("Will It Snow: ", weather.Forecast.ForecastDay[0].Hour[i].WillItSnow)
		blue("UV: ", weather.Forecast.ForecastDay[0].Hour[i].UV)
		if i < 23 {
			red("\n-------------------------------------Next Hour Forecast-------------------------------------")
		} else {
			red("\n------------------------------End of Forecast------------------------------")
		}
	}
}
