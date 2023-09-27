package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func getAddress() string {
	port := 8090
	hostname := os.Getenv("API_HOST")
	if hostname == "" {
		hostname = "localhost"
	}

	return fmt.Sprintf("%s:%d", hostname, port)
}

func startAPI(database *sql.DB) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	AddApiEndpoints(router, database)
	address := getAddress()
	router.Run(address)
}

func main() {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	arguments := os.Args[1:]
	if len(arguments) > 0 {
		HandleCommand(database, arguments)
	} else {
		startAPI(database)
	}
}
