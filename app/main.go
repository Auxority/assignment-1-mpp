package main

import (
	"fmt"
	"os"

	"mpp/command"
	"mpp/database"

	"github.com/gin-gonic/gin"
)

func getAddress() string {
	port := 8090
	database.QueryDatabase()
	hostname := os.Getenv("API_HOST")
	if hostname == "" {
		hostname = "localhost"
	}

	return fmt.Sprintf("%s:%d", hostname, port)
}

func startAPI() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	AddApiEndpoints(router)
	address := getAddress()
	router.Run(address)
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) > 0 {
		command.HandleCommand(arguments)
		return
	}

	go startAPI()
}
