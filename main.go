package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const WEATHER_KEY string = "92eb03e29fbeaac2ba48174fb79d8f27"
const WEATHER_API string = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s"

func main() {
	_res := "oui"
	for _res == "oui" {
		var cityName string
		fmt.Println("Température pour la ville : ")
		fmt.Scanln(&cityName)

		formattedUrl := fmt.Sprintf(WEATHER_API, cityName, WEATHER_KEY)
		response, err := http.Get(formattedUrl)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var weatherData WeatherData
		json.Unmarshal(data, &weatherData)
		fmt.Println(weatherData.tempToCelsius("Temp"))
		fmt.Println("Continuer ?")
		fmt.Scanln(&_res)
	}
}
