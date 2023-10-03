package command

import (
	"database/sql"
	"flag"
	"fmt"
	"mpp/database"
	"mpp/error_util"
	"mpp/types"
)

func createDetailsCommand() (*flag.FlagSet, *string) {
	name := "details"
	command := CreateNewCommand(&name)
	imdbIdParameter := CreateImdbIdParameter(command)

	return command, imdbIdParameter
}

func ShowDetails(id *string) {
	sql := fmt.Sprintf(`
		SELECT IMDb_id, Title, Rating, Year
		FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	database.QueryDatabase(&sql, showDetailsRow)
}

func showDetailsRow(rows *sql.Rows) {
	movie := getDetailsRow(rows)
	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear)
}

func getDetailsRow(rows *sql.Rows) *types.Movie {
	var movie types.Movie
	err := rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
	error_util.CheckError(err)
	return &movie
}
