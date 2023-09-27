package main

import (
	"database/sql"
	"fmt"
)

func ShowTitles(database *sql.DB) {
	sql := "SELECT Title FROM movies;"

	QueryDatabase(database, &sql, showTitleRow)
}

func showTitleRow(rows *sql.Rows) {
	var title string
	err := rows.Scan(&title)
	CheckError(err)
	fmt.Println(title)
}
