package main

import (
	"database/sql"
	"fmt"
)

func ShowTitles() {
	sql := "SELECT Title FROM movies;"

	QueryDatabase(&sql, showTitleRow)
}

func showTitleRow(rows *sql.Rows) {
	var title string
	err := rows.Scan(&title)
	CheckError(err)
	fmt.Println(title)
}
