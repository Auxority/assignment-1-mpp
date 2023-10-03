package database

import (
	"database/sql"
	"mpp/error_util"

	_ "github.com/mattn/go-sqlite3"
)

func CloseMoviesDatabase(database *sql.DB) {
	database.Close()
}

func ExecDatabase(sql *string) {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	_, err := database.Exec(*sql)
	error_util.CheckError(err)
}

func OpenMoviesDatabase() *sql.DB {
	fileName := "./movies.db"
	return openDatabase(&fileName)
}

func QueryDatabase(sql *string, nextRowFunc func(rows *sql.Rows)) {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	rows, err := database.Query(*sql)
	error_util.CheckError(err)
	defer closeRows(rows)

	for rows.Next() {
		nextRowFunc(rows)
	}

	defer rows.Close()
}

func openDatabase(fileName *string) *sql.DB {
	driverName := "sqlite3"
	database, err := sql.Open(driverName, *fileName)
	error_util.CheckError(err)
	return database
}

// TODO: Make private method
func closeRows(rows *sql.Rows) {
	rows.Close()
}
