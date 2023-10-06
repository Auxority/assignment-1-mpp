package command

import (
	"flag"
	"fmt"
	"mpp/api/database"
	"mpp/api/types"
)

func AddAndShowMovie(
	imdbId *string,
	title *string,
	imdbRating *float64,
	releaseYear *int,
) error {
	movie := types.Movie{IMDbId: imdbId, Title: title, IMDbRating: imdbRating, ReleaseYear: releaseYear}

	err := AddMovie(&movie)
	if err != nil {
		return fmt.Errorf("AddAndShowMovie: %w", err)
	}

	err = ShowMovieDetails(movie.IMDbId)
	if err != nil {
		return fmt.Errorf("AddAndShowMovie: %w", err)
	}

	return nil
}

func AddMovie(movie *types.Movie) error {
	sql := `
		INSERT INTO movies (IMDb_id, Title, Rating, Year)
		VALUES (?, ?, ?, ?);
	`
	err := database.ExecDatabase(&sql, *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear)
	if err != nil {
		return fmt.Errorf("AddMovie: %w", err)
	}

	return nil
}

func createAddCommand() (*flag.FlagSet, *string, *string, *float64, *int) {
	name := "add"
	command := CreateNewCommand(&name)

	imdbId := CreateImdbIdParameter(command)
	title := command.String("title", "Carmencita", "The movie's or series' title")
	imdbRating := command.Float64("rating", 5.7, "The movie's or series' rating on IMDb")
	releaseYear := command.Int("year", 1894, "The movie's or series' year of release")

	return command, imdbId, title, imdbRating, releaseYear
}
