package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}

	fmt.Println("Enter your city name")
	fmt.Scanln(&input)
	if err != nil {
		log.Fatal("An Error occured :", err)
	}
	if input == "" {
		// Empty Input
		input = "London" // Defaulting to London
		fmt.Println("No City name Provided. Choose Any from the below list")
		fmt.Println("1.Varanasi\n2.Paris\n3.New York\n4.Moscow")
		var option string
		fmt.Scanln(&option)
		if option == "" {
			fmt.Println("No Options chosen. You lazy Bastard :/ :/ ........ \nDefaulting to London ........")
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
				fmt.Println("My Chanpagne Lover!! Please choose a correct option next time. Will you? :/ \n Defaulting to London !! :/ .....")
				input = "London"
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
	heading := fmt.Sprintf("\n------------------------------More Information on %s------------------------------\n", input)
	fmt.Println(heading)
	fmt.Println("Location: ", weather.Location.Name)
	fmt.Println("Region: ", weather.Location.Region)
	fmt.Println("Country: ", weather.Location.Country)
	fmt.Println("Latitude: ", weather.Location.Lat)
	fmt.Println("Longitude: ", weather.Location.Lon)
	fmt.Println("Temperature: ", weather.Current.TempC)
	fmt.Println("Wind Speed: ", weather.Current.WindMph)
	fmt.Println("Wind Degree: ", weather.Current.WindDegree)
	fmt.Println("Pressure: ", weather.Current.Pressure)
	fmt.Println("Precipitation: ", weather.Current.PrecipMm)
	fmt.Println("Cloud: ", weather.Current.Cloud)
	fmt.Println("Feels Like: ", weather.Current.FeelslikeC)
	fmt.Println("Heat Index: ", weather.Current.HeatindexC)
	fmt.Println("Condition: ", weather.Current.Condition.Text)
	fmt.Println("Dew Point: ", weather.Current.DewPoint)
	fmt.Println("\n------------------------------Forecast------------------------------\n")
	fmt.Println("Max Daily Temperature", weather.Forecast.ForecastDay[0].Day.MaxtempC)
	fmt.Println("Min Daily Temperature ", weather.Forecast.ForecastDay[0].Day.MintempC)
	fmt.Println("Average Daily Temperature ", weather.Forecast.ForecastDay[0].Day.AvgtempC)
	fmt.Println("Max Wind Speed ", weather.Forecast.ForecastDay[0].Day.MaxwindMph)
	fmt.Println("Total Precipitation ", weather.Forecast.ForecastDay[0].Day.TotalprecipMm)
	fmt.Println("Total Snow ", weather.Forecast.ForecastDay[0].Day.TotalsnowCm)
	fmt.Println("Average Humidity ", weather.Forecast.ForecastDay[0].Day.AvgHumidity)
	fmt.Println("Condition ", weather.Forecast.ForecastDay[0].Day.Condition.Text)
	fmt.Println("UV ", weather.Forecast.ForecastDay[0].Day.UV)
	fmt.Println("\n------------------------------Astro------------------------------\n")
	fmt.Println("Sunrise ", weather.Forecast.ForecastDay[0].Astro.Sunrise)
	fmt.Println("Sunset ", weather.Forecast.ForecastDay[0].Astro.Sunset)
	fmt.Println("Moonrise ", weather.Forecast.ForecastDay[0].Astro.Moonrise)
	fmt.Println("Moonset ", weather.Forecast.ForecastDay[0].Astro.Moonset)
	fmt.Println("\n------------------------------Hourly Forecast------------------------------\n")
	for i := 0; i < 24; i++ {
		fmt.Println("Time: ", weather.Forecast.ForecastDay[0].Hour[i].Time)
		fmt.Println("Condition: ", weather.Forecast.ForecastDay[0].Hour[i].Condition.Text)
		fmt.Println("Feels Like: ", weather.Forecast.ForecastDay[0].Hour[i].FeelslikeC)
		fmt.Println("Humidity: ", weather.Forecast.ForecastDay[0].Hour[i].Humidity)
		fmt.Println("Cloud: ", weather.Forecast.ForecastDay[0].Hour[i].Cloud)
		fmt.Println("Will It Rain: ", weather.Forecast.ForecastDay[0].Hour[i].WillItRain)
		fmt.Println("Wind Speed: ", weather.Forecast.ForecastDay[0].Hour[i].WindMph)
		fmt.Println("Will It Snow: ", weather.Forecast.ForecastDay[0].Hour[i].WillItSnow)
		fmt.Println("UV: ", weather.Forecast.ForecastDay[0].Hour[i].UV)
		if i < 23 {
			fmt.Println("\n-------------------------------------Next Hour Forecast-------------------------------------")
		} else {
			fmt.Println("\n------------------------------End of Forecast------------------------------")
		}
	}
}
