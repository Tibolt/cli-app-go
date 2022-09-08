package util

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type DiceCommand struct {
	fs *flag.FlagSet
	dice string
	numRoll int
}

func NewDiceCommand() *DiceCommand {
	dc  := &DiceCommand {
		fs: flag.NewFlagSet("dice", flag.ContinueOnError),
	}
	dc.fs.StringVar(&dc.dice, "d", "d6", "Type of dice to roll. Pick dX where X is int.")
	dc.fs.IntVar(&dc.numRoll, "n", 1, "Number of dice to roll.")
	return dc
}

func (d *DiceCommand) Name() string {
	return d.fs.Name()
}


func (d *DiceCommand) Init(args []string) error {
	return d.fs.Parse(args)
}

func (d *DiceCommand) Run() error {
	RollMain(&d.dice, &d.numRoll)
	return nil
}

func RollDice(dice *string, number *int) []int{
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
	rand.Seed(time.Now().UTC().UnixNano())
	matched, _ := regexp.Match("d\\d+", []byte(*dice))

	if matched == true {
		rolls := RollDice(dice, numRoll)
		printDice(rolls)

	} else {
		fmt.Print("Wrong type of dice!")
	}
}
