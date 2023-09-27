package main

import (
	"database/sql"
	"flag"
	"fmt"
)

func createDetailsCommand() (*flag.FlagSet, *string) {
	name := "details"
	command := CreateNewCommand(&name)
	imdbIdParameter := CreateImdbIdParameter(command)

	return command, imdbIdParameter
}

func ShowDetails(database *sql.DB, id *string) {
	sql := fmt.Sprintf(`
		SELECT IMDb_id, Title, Rating, Year
		FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	QueryDatabase(database, &sql, showDetailsRow)
}

func showDetailsRow(rows *sql.Rows) {
	var movie Movie
	getDetailsRow(rows, &movie)
	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear)
}

func getDetailsRow(rows *sql.Rows, movie *Movie) {
	err := rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
	CheckError(err)
}
