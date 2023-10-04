package main

import (
	"os"

	"mpp/command"
	"mpp/error_util"
	"mpp/router"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) > 0 {
		err := command.HandleCommand(arguments)
		error_util.CheckError(err)
		return
	}

	router.StartAPI()
}
