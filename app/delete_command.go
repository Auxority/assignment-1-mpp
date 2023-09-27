package main

import (
	"database/sql"
	"flag"
	"fmt"
)

func createDeleteCommand() (*flag.FlagSet, *string) {
	name := "delete"
	command := CreateNewCommand(&name)
	imdbIdParameter := CreateImdbIdParameter(command)

	return command, imdbIdParameter
}

func DeleteMovie(database *sql.DB, id *string) {
	sql := fmt.Sprintf(`
		DELETE FROM movies
		WHERE IMDb_id='%s';
	`, *id)
	ExecDatabase(database, &sql)
}

func DeleteMovieCommand(database *sql.DB, id *string) {
	DeleteMovie(database, id)
	fmt.Println("Movie deleted")
}
