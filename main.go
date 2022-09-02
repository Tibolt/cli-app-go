package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"example/app/util"
)

func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	dice := flag.String("d", "d6", "Type of dice to roll. Pick dX where X is int.")
	numRoll := flag.Int("n", 1, "Number of dice to roll.")
	city := flag.String("c", "Warsaw", "Name of city")
	flag.Parse()

	rollMain(dice, numRoll)
	util.GetWeather(city)

}

func rollDice(dice *string, number *int) []int{
	var rolls []int

	diceSides := (*dice)[1:]
	d, err := strconv.Atoi(diceSides)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < *number; i++{
		rolls = append(rolls, rand.Intn(d) + 1)
	}
	return rolls
}

func printDice(rolls []int){
	for i, dice := range rolls{
		fmt.Printf("Dice %d rolled %d\n", i, dice)
	}
}

func rollMain(dice *string, numRoll *int){
	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched == true {
		rolls := rollDice(dice, numRoll)
		printDice(rolls)

	} else {
		fmt.Print("Wrong type of dice!")
	}
}
