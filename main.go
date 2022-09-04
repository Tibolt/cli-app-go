package main

import (
	"errors"
	"example/app/util"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type GreetCommand struct {
	fs *flag.FlagSet
	name string
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func NewGreetCommand() *GreetCommand {
	gc := &GreetCommand{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}
	gc.fs.StringVar(&gc.name, "name", "World", "name of the person to be greeted")
	return gc
}

func (g *GreetCommand) Name() string {
	return g.fs.Name()
}

func (g *GreetCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *GreetCommand) Run() error {
	fmt.Println("Hello", g.name, "!")
	return nil
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass command")
	}

	cmds := []Runner{
		NewGreetCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("Unknown subcommand: %s ", subcommand)
}

func main() {
	err := root(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Avaible commands: greet\n")
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

