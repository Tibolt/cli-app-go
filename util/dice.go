package util

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
)

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
		fmt.Printf("Roll number %d was: %d\n", i, dice)
	}
}

func RollMain(dice *string, numRoll *int){
	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched == true {
		rolls := rollDice(dice, numRoll)
		printDice(rolls)

	} else {
		fmt.Print("Wrong type of dice!")
	}
}
