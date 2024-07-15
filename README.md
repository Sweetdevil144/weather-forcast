## Readme.md

### Overview
This Go-based command-line application provides detailed weather forecasts. Utilizing real-time data from the Weather API, the app displays current weather conditions, daily and hourly forecasts, and astronomical data for specified locations.

### Features
- **Current Weather Data:** Displays temperature, wind speed, humidity, and more for the user's selected location.
- **Daily Forecast:** Offers a complete outlook, including temperature variations, precipitation, and UV index.
- **Hourly Forecast:** Details hourly weather conditions for the next 24 hours.
- **Astronomical Data:** Shows times for sunrise, sunset, moonrise, and moonset.
- **Interactive Input:** Users can enter a city name or select from a predefined list of cities if no input is provided.

### Installation

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/your-username/weather-forecast.git
   ```
2. **Navigate to the Directory:**
   ```sh
   cd weather-forecast
   ```
3. **Install Dependencies:**
   ```sh
   go mod tidy
   ```

### Configuration

- The application requires an API key from Weather API, which should be placed in a `.env` file:
  ```env
  API_KEY=your_api_key_here
  ```

### Usage

- **Run the Application:**
  ```sh
  go run main.go [city name] (optional)
  ```
  Optionally, you can run the application without the city name to interactively input or select a city.
  
- **Example Command:**
  ```sh
  go run main.go London
  ```

### Output Example
Upon running the application, you will see output similar to this:

```
------------------------------More Information on LONDON------------------------------
Location: London
Region: London
Country: United Kingdom
Latitude: 51.5074
Longitude: -0.1278
Temperature: 15.0°C
Wind Speed: 10 mph
..........
```

### Contributing

1. **Fork the Repository:** Start by forking the repository on GitHub.
2. **Create a Branch:** Create a new branch for your modifications (`git checkout -b new-feature`).
3. **Commit Changes:** After making your changes, commit them (`git commit -am 'Added some feature'`).
4. **Push to the Branch:** Push your branch to GitHub (`git push origin new-feature`).
5. **Open a Pull Request:** Go to the repository page on GitHub and click the “Compare & pull request” button.

### License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
