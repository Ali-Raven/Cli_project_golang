package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
)

func main() {
	myFig := figure.NewColorFigure("CLI app", "", "blue", true)
	myFig.Print()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run <program> arg1 arg2")
		return
	}

	firstArg := os.Args[1]

	switch firstArg {
	case "--add":
		{
			AddTasks()
		}
	case "--edit":
		{
			EditTasks()
		}
	case "--del":
		{
			DeleteTasks()
		}
	default: {
		fmt.Println("use one of our flag : --add , --edit , --del")
		return
	}
	}
	fmt.Println(firstArg)
}
