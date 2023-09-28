package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CloseMoviesDatabase(database *sql.DB) {
	database.Close()
}

func ExecDatabase(sql *string) {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	_, err := database.Exec(*sql)
	CheckError(err)
}

func OpenMoviesDatabase() *sql.DB {
	fileName := "./movies.db"
	return openDatabase(&fileName)
}

func QueryDatabase(sql *string, nextRowFunc func(rows *sql.Rows)) {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	rows, err := database.Query(*sql)
	CheckError(err)
	defer closeRows(rows)

	for rows.Next() {
		nextRowFunc(rows)
	}

	defer rows.Close()
}

func openDatabase(fileName *string) *sql.DB {
	driverName := "sqlite3"
	database, err := sql.Open(driverName, *fileName)
	CheckError(err)
	return database
}

func closeRows(rows *sql.Rows) {
	rows.Close()
}
