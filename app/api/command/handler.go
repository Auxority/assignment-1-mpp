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
	addCommand, addImdbId, addTitle, addImdbRating, addYear, addPlot := createAddCommand()
	detailsCommand, detailsImdbId := createDetailsCommand()
	deleteCommand, deleteImdbId := createDeleteCommand()

	switch arguments[0] {
	case "add":
		parseArguments(addCommand, &arguments)
		return AddAndShowMovie(addImdbId, addTitle, addImdbRating, addYear, addPlot)
	case "list":
		return ShowMovieList()
	case "details":
		parseArguments(detailsCommand, &arguments)
		return ShowMovieDetails(detailsImdbId)
	case "delete":
		parseArguments(deleteCommand, &arguments)
		return ShowMovieDeletion(deleteImdbId)
	case "summaries":
		return ShowMovieSummaries()
	default:
		return fmt.Errorf("HandleCommand: unable to find the provided '%q' command", arguments[0])
	}
}
