package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Data struct{
	Weather []Weather `json:"weather"`
	Info Info	`json:"main"`
}
type Weather struct{
	Main string `json:"main"`
}
type Info struct{
	Temp float32 `json:"temp"`
	Humidity int `json:"humidity"`
} 

func GetWeather(city *string){
	url := 	fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", *city, os.Getenv("WEATHER_API"))
	response, err := http.Get(url)
	if err != nil {
        fmt.Print(err.Error())
		os.Exit(1)
    }
	defer response.Body.Close()

	responeData, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}

	var weatherObject Data
	json.Unmarshal(responeData, &weatherObject)
	fmt.Printf("Temprature in %s is: %f\n", *city, kelvinToCelcius(weatherObject.Info.Temp))
	fmt.Printf("Humidity: %d\n", weatherObject.Info.Humidity)

	for _, obj := range weatherObject.Weather{
		fmt.Printf(obj.Main)
	}
}

func kelvinToCelcius(temp float32) float32{
	var k float32 = 273.15
	return temp - k
}
