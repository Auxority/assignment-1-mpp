package main

import (
	"os"

	"mpp/command"
	"mpp/router"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) > 0 {
		command.HandleCommand(arguments)
		return
	}

	router.StartAPI()
}
