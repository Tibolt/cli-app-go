package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WeatherCommand struct {
	fs *flag.FlagSet
	city string
}

func NewWeatherCommand() *WeatherCommand {
	wc := &WeatherCommand{
		fs: flag.NewFlagSet("weather", flag.ContinueOnError),
	}
	wc.fs.StringVar(&wc.city, "city", "Warsaw", "Name of the city")
	return wc
}

func (w *WeatherCommand) Name() string {
	return w.fs.Name()
}

func (w *WeatherCommand) Init(args []string) error {
	return w.fs.Parse(args)
}

func (w *WeatherCommand) Run() error {
	GetWeather(&w.city)
	return nil
}

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
	fmt.Printf("Temperature in %s is: %f\n", *city, kelvinToCelcius(weatherObject.Info.Temp))
	fmt.Printf("Humidity: %d\n", weatherObject.Info.Humidity)

	for _, obj := range weatherObject.Weather{
		fmt.Printf(obj.Main)
	}
}

func kelvinToCelcius(temp float32) float32{
	var k float32 = 273.15
	return temp - k
}
