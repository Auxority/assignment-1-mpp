package command

import (
	"mpp/omdb"
)

func printMovieSummary(id *string) {
	var summary omdb.MovieSummary
	go omdb.GetMovieSummary(id, &summary)

}

func PrintMovieSummaries() {
	imdbId := "tt0034583"
	printMovieSummary(&imdbId)
}
