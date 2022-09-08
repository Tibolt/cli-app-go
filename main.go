package main

import (
	"example/app/util"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	err := util.Root(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Avaible commands: greet, dice, weather\n")
		os.Exit(1)
	}
}

func oldMain() {
	rand.Seed(time.Now().UTC().UnixNano())
	dice := flag.String("d", "d6", "Type of dice to roll. Pick dX where X is int.")
	numRoll := flag.Int("n", 1, "Number of dice to roll.")
	city := flag.String("c", "Warsaw", "Name of city")
	flag.Parse()

	util.RollMain(dice, numRoll)
	util.GetWeather(city)
}

