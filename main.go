package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}

	fmt.Println("Your API Key is:", apiKey)
	input, err := fmt.Scanln("Enter your city name")
	if err != nil {
		log.Fatal("An Error occured :\n",err)
	}

	query := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", apiKey, input)
	res, err := http.Get(query)
	
	if err != nil {
		log.Fatal("An Error occured :\n",err)
	} else {
		fmt.Println(res)
	}
}
