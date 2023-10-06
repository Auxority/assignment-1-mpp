package command

import (
	"fmt"
	"mpp/api/database"
	"mpp/api/types"
)

func ShowMovieList() error {
	movies, err := GetMovieList()
	if err != nil {
		return fmt.Errorf("ShowMovieList: %w", err)
	}

	for _, movie := range movies {
		fmt.Println(*movie.Title)
	}

	return nil
}

func GetMovieList() ([]*types.Movie, error) {
	sql := "SELECT IMDb_id, Title, Rating, Year, Plot_summary FROM movies;"
	rows, err := database.QueryDatabase(&sql, getMovieFromRow)

	if err != nil {
		return nil, fmt.Errorf("GetMovieList: %w", err)
	}

	var movies []*types.Movie
	for _, row := range rows {
		movie := (*row).(*types.Movie)
		movies = append(movies, movie)
	}

	return movies, nil
}
