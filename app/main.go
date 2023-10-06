package main

import (
	"os"

	"mpp/api/command"
	"mpp/api/router"
	"mpp/error_util"
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
