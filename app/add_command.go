package main

import (
	"flag"
	"fmt"
)

func createAddCommand() (*flag.FlagSet, *string, *string, *float64, *int) {
	name := "add"
	command := CreateNewCommand(&name)

	imdbId := CreateImdbIdParameter(command)
	title := command.String("title", "Carmencita", "The movie's or series' title")
	imdbRating := command.Float64("rating", 5.7, "The movie's or series' rating on IMDb")
	releaseYear := command.Int("year", 1894, "The movie's or series' year of release")

	return command, imdbId, title, imdbRating, releaseYear
}

func AddMovie(
	movie *Movie,
) {
	sql := fmt.Sprintf(`
		INSERT INTO movies (IMDb_id, Title, Rating, Year)
		VALUES ('%s', '%s', %.1f, %d);
	`, *movie.IMDbId, *movie.Title, *movie.IMDbRating, *movie.ReleaseYear)

	ExecDatabase(&sql)
}

func AddMovieCommand(
	imdbId *string,
	title *string,
	imdbRating *float64,
	releaseYear *int,
) {
	movie := Movie{IMDbId: imdbId, Title: title, IMDbRating: imdbRating, ReleaseYear: releaseYear}
	AddMovie(&movie)
	ShowDetails(movie.IMDbId)
}
