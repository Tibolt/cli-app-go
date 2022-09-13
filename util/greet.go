package util

import (
	"flag"
	"fmt"
)


type GreetCommand struct {
	fs *flag.FlagSet
	name string
}

func NewGreetCommand() *GreetCommand {
	gc := &GreetCommand{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}
	gc.fs.StringVar(&gc.name, "name", "World", "Name of the person to be greeted")
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


