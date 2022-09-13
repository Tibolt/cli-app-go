package main

import (
	"example/app/util"
	"fmt"
	"os"
)

func main() {
	err := util.Root(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Avaible commands: help, greet, dice, weather\n")
		os.Exit(1)
	}
}

