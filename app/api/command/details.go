package command

import (
	"database/sql"
	"flag"
	"fmt"
	"mpp/api/database"
	"mpp/api/types"
)

func ShowMovieDetails(id *string) error {
	movie, err := GetMovieDetails(id)
	if err != nil {
		return fmt.Errorf("ShowDetails: %w", err)
	}

	fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear)

	return nil
}

func GetMovieDetails(id *string) (*types.Movie, error) {
	sql := `
		SELECT IMDb_id, Title, Rating, Year, Plot_summary
		FROM movies
		WHERE IMDb_id = ?;
	`

	movies, err := database.QueryDatabase(&sql, getMovieFromRow, *id)
	if err != nil {
		return nil, fmt.Errorf("GetMovieDetails: %w", err)
	}

	movie := (*movies[0]).(*types.Movie)

	return movie, nil
}

func createDetailsCommand() (*flag.FlagSet, *string) {
	name := "details"
	command := CreateNewCommand(&name)
	imdbIdParameter := CreateImdbIdParameter(command)

	return command, imdbIdParameter
}

func getMovieFromRow(rows *sql.Rows) (any, error) {
	var movie types.Movie
	err := rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear, &movie.Plot_summary)
	if err != nil {
		return nil, fmt.Errorf("getMovieFromRow: failed to scan rows: %w", err)
	}

	return &movie, nil
}
