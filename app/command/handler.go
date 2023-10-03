package command

import (
	"flag"
	"fmt"
)

func parseArguments(command *flag.FlagSet, arguments *[]string) {
	commandArguments := (*arguments)[1:]
	command.Parse(commandArguments)
}

func HandleCommand(arguments []string) error {
	addCommand, addImdbId, addTitle, addImdbRating, addYear := createAddCommand()
	detailsCommand, detailsImdbId := createDetailsCommand()
	deleteCommand, deleteImdbId := createDeleteCommand()

	switch arguments[0] {
	case "add":
		parseArguments(addCommand, &arguments)
		AddMovieCommand(addImdbId, addTitle, addImdbRating, addYear)
	case "list":
		ShowTitles()
	case "details":
		parseArguments(detailsCommand, &arguments)
		ShowDetails(detailsImdbId)
	case "delete":
		parseArguments(deleteCommand, &arguments)
		DeleteMovieCommand(deleteImdbId)
	case "summaries":
		fmt.Println("aaaa")
		PrintMovieSummaries()
	default:
		return fmt.Errorf("unable to find the provided '%q' command", arguments[0])
	}

	return nil
}
