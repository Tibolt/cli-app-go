package util

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass command")
	}

	cmds := []Runner{
		NewGreetCommand(),
		NewDiceCommand(),
		NewWeatherCommand(),
		NewHelpCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])

			//Fix for executing command with -h flag
			if len(os.Args[2:]) != 0{
				matched, _ := regexp.Match("-h", []byte(os.Args[2]))
				if matched {
					return nil	
				}
			}
			return cmd.Run()
		}
	}
	return fmt.Errorf("Unknown subcommand: %s ", subcommand)
}

