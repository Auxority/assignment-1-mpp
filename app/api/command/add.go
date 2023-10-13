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
	plot *string,
) error {
	movie := types.Movie{IMDbId: imdbId, Title: title, IMDbRating: imdbRating, ReleaseYear: releaseYear, Plot_summary: plot}

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
	sql := `INSERT INTO movies (IMDb_id, Title, Rating, Year, Plot_summary) VALUES (?, ?, ?, ?, ?);`

	err := database.ExecDatabase(&sql, *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear, *movie.Plot_summary)
	if err != nil {
		return fmt.Errorf("AddMovie: %w", err)
	}

	return nil
}

func createAddCommand() (*flag.FlagSet, *string, *string, *float64, *int, *string) {
	name := "add"
	command := CreateNewCommand(&name)

	imdbId := CreateImdbIdParameter(command)
	title := command.String("title", "Boing", "The movie's or series' title")
	imdbRating := command.Float64("rating", 9.9, "The movie's or series' rating on IMDb")
	releaseYear := command.Int("year", 2023, "The movie's or series' year of release")
	plot := command.String("plot", "This has a great plot.", "The movie's or series' plot summary")

	return command, imdbId, title, imdbRating, releaseYear, plot
}
