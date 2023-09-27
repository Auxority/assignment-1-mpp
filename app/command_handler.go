package main

import (
	"database/sql"
	"flag"
	"fmt"
)

func parseArguments(command *flag.FlagSet, arguments *[]string) {
	commandArguments := (*arguments)[1:]
	command.Parse(commandArguments)
}

func HandleCommand(database *sql.DB, arguments []string) error {
	addCommand, addImdbId, addTitle, addImdbRating, addYear := createAddCommand()
	detailsCommand, detailsImdbId := createDetailsCommand()
	deleteCommand, deleteImdbId := createDeleteCommand()

	switch arguments[0] {
	case "add":
		parseArguments(addCommand, &arguments)
		AddMovieCommand(database, addImdbId, addTitle, addImdbRating, addYear)
	case "list":
		ShowTitles(database)
	case "details":
		parseArguments(detailsCommand, &arguments)
		ShowDetails(database, detailsImdbId)
	case "delete":
		parseArguments(deleteCommand, &arguments)
		DeleteMovieCommand(database, deleteImdbId)
	default:
		return fmt.Errorf("unable to find the provided '%q' command", arguments[0])
	}

	return nil
}
