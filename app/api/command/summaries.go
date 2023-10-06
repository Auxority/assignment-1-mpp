package command

import (
	"fmt"
	"mpp/api/database"
	"mpp/api/omdb"
	"mpp/api/types"
	"sync"
)

func ShowMovieSummaries() error {
	movies, err := GetMovieList()
	if err != nil {
		return fmt.Errorf("ShowMovieSummaries: %w", err)
	}

	var waitGroup sync.WaitGroup
	errChannel := make(chan error, len(movies))
	maxGoRoutines := 10
	limitChannel := make(chan struct{}, maxGoRoutines)

	for _, movie := range movies {
		waitGroup.Add(1)
		limitChannel <- struct{}{}
		go showMovieSummary(&waitGroup, movie, &errChannel, &limitChannel)
	}

	waitGroup.Wait()
	close(errChannel)

	for err := range errChannel {
		if err != nil {
			return fmt.Errorf("ShowMovieSummaries: %w", err)
		}
	}

	fmt.Println("Summaries added")

	return nil
}

func showMovieSummary(waitGroup *sync.WaitGroup, movie *types.Movie, errChannel *chan error, limitConcurrency *chan struct{}) {
	defer waitGroup.Done()
	err := addAndShowMovieSummary(movie)
	if err != nil {
		*errChannel <- err
	}
	<-*limitConcurrency
}

func addAndShowMovieSummary(movie *types.Movie) error {
	var summary types.MovieSummary
	err := omdb.GetMovieSummary(movie.IMDbId, &summary)
	if err != nil {
		return fmt.Errorf("addAndShowMovieSummary: %w", err)
	}

	sql := `
			UPDATE movies
			SET Plot_summary = ?
			WHERE imdb_id = ?
		`
	err = database.ExecDatabase(&sql, summary.Plot, *movie.IMDbId)
	if err != nil {
		return fmt.Errorf("addAndShowMovieSummary: %w", err)
	}

	return nil
}
