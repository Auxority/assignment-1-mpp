package command

import (
	"flag"
	"fmt"
	"mpp/database"
)

func createDeleteCommand() (*flag.FlagSet, *string) {
	name := "delete"
	command := CreateNewCommand(&name)
	imdbIdParameter := CreateImdbIdParameter(command)

	return command, imdbIdParameter
}

func DeleteMovie(id *string) error {
	sql := `
		DELETE FROM movies
		WHERE IMDb_id = ?;
	`

	err := database.ExecDatabase(&sql, *id)
	if err != nil {
		return fmt.Errorf("DeleteMovie: %w", err)
	}

	return nil
}

func ShowMovieDeletion(id *string) error {
	err := DeleteMovie(id)
	if err != nil {
		return fmt.Errorf("ShowMovieDeletion: %w", err)
	}

	fmt.Println("Movie deleted")

	return nil
}
