package util

import (
	"flag"
	"fmt"
)

type HelpCommand struct {
	fs *flag.FlagSet
}

func NewHelpCommand() *HelpCommand {
	gc := &HelpCommand{
		fs: flag.NewFlagSet("help", flag.ContinueOnError),
	}
	return gc
}

func (g *HelpCommand) Name() string {
	return g.fs.Name()
}

func (g *HelpCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *HelpCommand) Run() error {
	fmt.Printf("CLI app to get current weather info and roll dice\n")
	fmt.Printf("Avaible subcommands: help, greet, weather, dice\n")
	fmt.Printf("To get more info about subcommand type <command> -h\n")
	return nil
}


