package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DRIVER_NAME  string = "sqlite3"
	DB_FILE_NAME string = "./data/movies.db"
)

func OpenDatabase() (*sql.DB, error) {
	database, err := sql.Open(DRIVER_NAME, DB_FILE_NAME)
	if err != nil {
		return nil, fmt.Errorf("InitDatabase: failed to open database connection: %w", err)
	}
	return database, nil
}

func ExecDatabase(sql *string, args ...any) error {
	database, err := OpenDatabase()
	if err != nil {
		return fmt.Errorf("ExecDatabase: %w", err)
	}

	defer database.Close()

	_, err = database.Exec(*sql, args...)
	if err != nil {
		return fmt.Errorf("ExecDatabase: failed to execute SQL: %w", err)
	}

	return nil
}

func QueryDatabase(sql *string, nextRowFunc func(rows *sql.Rows) (any, error), args ...any) ([]*any, error) {
	database, err := OpenDatabase()
	if err != nil {
		return nil, fmt.Errorf("QueryDatabase: %w", err)
	}
	defer database.Close()

	rows, err := database.Query(*sql, args...)
	if err != nil {
		fmt.Println(*sql)
		fmt.Println(args...)
		return nil, fmt.Errorf("QueryDatabase: failed to query database: %w", err)
	}
	defer rows.Close()

	var results []*any
	for rows.Next() {
		result, err := nextRowFunc(rows)
		if err != nil {
			return nil, fmt.Errorf("QueryDatabase: failed to run next row function: %w", err)
		}
		results = append(results, &result)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("QueryDatabase: no results found")
	}

	return results, nil
}
