package command

import (
	"fmt"
	"mpp/api/database"
	"mpp/api/types"
)

func ShowMovieList(page int, limit int) error {
	movies, err := GetMovieList(page, limit)
	if err != nil {
		return fmt.Errorf("ShowMovieList: %w", err)
	}

	for _, movie := range movies {
		fmt.Println(*movie.Title)
	}

	return nil
}

func GetMovieList(page int, limit int) ([]*types.Movie, error) {
	sql := `
		SELECT IMDb_id, Title, Rating, Year, Plot_summary
		FROM movies
		ORDER BY Year DESC
		LIMIT $1
		OFFSET $2
	`
	offset := page * limit
	rows, err := database.QueryDatabase(&sql, getMovieFromRow, limit, offset)

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
