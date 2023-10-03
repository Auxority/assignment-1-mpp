package command

import (
	"database/sql"
	"fmt"
	"mpp/database"
	"mpp/error_util"
)

func ShowTitles() {
	sql := "SELECT Title FROM movies;"

	database.QueryDatabase(&sql, showTitleRow)
}

func showTitleRow(rows *sql.Rows) {
	var title string
	err := rows.Scan(&title)
	error_util.CheckError(err)
	fmt.Println(title)
}
