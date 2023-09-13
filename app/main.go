package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func show_titles(db *sql.DB) {
	query := "SELECT Title FROM movies;"

	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var title string
		err = rows.Scan(&title)
		checkError(err)
		fmt.Println(title)
	}
}

func show_details(db *sql.DB, id *string) {
	query := fmt.Sprintf(`
		SELECT * FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var imdbId string
		var title string
		var imdbRating float64
		var releaseYear int
		err = rows.Scan(&imdbId, &title, &imdbRating, &releaseYear)
		checkError(err)
		fmt.Printf("IMDb id: %s\nTitle: %s\nRating: %.1f\nYear: %d\n", imdbId, title, imdbRating, releaseYear)
	}
}

func add_movie(
	db *sql.DB,
	id *string,
	title *string,
	releaseYear *int,
	imdbRating *float64,
) {
	query := fmt.Sprintf(`
		INSERT INTO movies (IMDb_id, Title, Rating, Year)
		VALUES ('%s', '%s', %.1f, %d);
	`, *id, *title, *imdbRating, *releaseYear)

	_, err := db.Exec(query)
	checkError(err)

	show_details(db, id)
}

func delete_movie(db *sql.DB, id *string) {
	query := fmt.Sprintf(`
		DELETE FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	_, err := db.Exec(query)
	checkError(err)

	fmt.Println("Movie deleted")
}

func main() {
	arguments := os.Args[1:]

	db, err := sql.Open("sqlite3", "./movies.db")
	checkError(err)
	defer db.Close()

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)

	addImdbId := addCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")
	addTitle := addCommand.String("title", "Carmencita", "The movie's or series' title")
	addImdbRating := addCommand.Float64("rating", 5.7, "The movie's or series' rating on IMDb")
	addYear := addCommand.Int("year", 1894, "The movie's or series' year of release")

	detailsCommand := flag.NewFlagSet("details", flag.ExitOnError)
	detailsImdbId := detailsCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteImdbId := deleteCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")

	switch arguments[0] {
	case "add":
		addCommand.Parse(arguments[1:])
		add_movie(db, addImdbId, addTitle, addYear, addImdbRating)
	case "list":
		show_titles(db)
	case "details":
		detailsCommand.Parse(arguments[1:])
		show_details(db, detailsImdbId)
	case "delete":
		deleteCommand.Parse(arguments[1:])
		delete_movie(db, deleteImdbId)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
