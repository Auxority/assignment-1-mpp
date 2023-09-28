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

func ShowDetails(id *string) {
	sql := fmt.Sprintf(`
		SELECT IMDb_id, Title, Rating, Year
		FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	QueryDatabase(&sql, showDetailsRow)
}

func showDetailsRow(rows *sql.Rows) {
	movie := getDetailsRow(rows)
	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear)
}

func getDetailsRow(rows *sql.Rows) *Movie {
	var movie Movie
	err := rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
	CheckError(err)
	return &movie
}
