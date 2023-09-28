package main

import (
	"flag"
	"fmt"
)

func createDeleteCommand() (*flag.FlagSet, *string) {
	name := "delete"
	command := CreateNewCommand(&name)
	imdbIdParameter := CreateImdbIdParameter(command)

	return command, imdbIdParameter
}

func DeleteMovie(id *string) {
	sql := fmt.Sprintf(`
		DELETE FROM movies
		WHERE IMDb_id='%s';
	`, *id)
	ExecDatabase(&sql)
}

func DeleteMovieCommand(id *string) {
	DeleteMovie(id)
	fmt.Println("Movie deleted")
}
