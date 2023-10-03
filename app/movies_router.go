package main

import (
	"fmt"
	"mpp/command"
	"mpp/database"
	"mpp/error_util"
	"mpp/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Find a better way that makes use of interfaces in the queryDatabase function to query the database.
// Since this would prevent a lot of code duplication.
func queryMovies(movies *[]*types.Movie) {
	moviesDatabase := database.OpenMoviesDatabase()
	defer database.CloseMoviesDatabase(moviesDatabase)

	rows, err := moviesDatabase.Query("SELECT IMDb_id, Title, Rating, Year FROM movies;")
	error_util.CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var movie types.Movie
		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
		*movies = append(*movies, &movie)
	}

	defer rows.Close()
}

// TODO: IDEM, ideally the commands would all return structs that are serialized to JSON.
// But this requires a refactor, while I just want to get everything working for now.
func queryMovie(id *string) *types.Movie {
	moviesDatabase := database.OpenMoviesDatabase()
	defer database.CloseMoviesDatabase(moviesDatabase)

	sql := fmt.Sprintf(`
		SELECT IMDb_id, Title, Rating, Year
		FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	rows, err := moviesDatabase.Query(sql)
	error_util.CheckError(err)
	defer rows.Close()

	var movie types.Movie
	for rows.Next() {
		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
	}

	defer rows.Close()

	return &movie
}

func movieListEndpoint(router *gin.Engine) {
	router.GET("/movies", func(context *gin.Context) {
		var movies []*types.Movie
		queryMovies(&movies)

		context.IndentedJSON(http.StatusOK, &movies)
	})
}

func movieDetailsEndpoint(router *gin.Engine) {
	router.GET("/movies/:id", func(context *gin.Context) {
		id := context.Param("id")
		movie := queryMovie(&id)

		if movie.IMDbId != nil {
			context.IndentedJSON(http.StatusOK, &movie)
		} else {
			context.Status(http.StatusNotFound)
		}
	})
}

func movieDeleteEndpoint(router *gin.Engine) {
	router.DELETE("/movies/:id", func(context *gin.Context) {
		id := context.Param("id")
		command.DeleteMovie(&id)
		context.Status(http.StatusNoContent)
	})
}

func movieAddEndpoint(router *gin.Engine) {
	router.POST("/movies", func(context *gin.Context) {
		var movie types.Movie

		err := context.BindJSON(&movie)
		error_util.CheckError(err)

		command.AddMovie(&movie)

		context.IndentedJSON(http.StatusOK, &movie)
	})
}

func AddApiEndpoints(router *gin.Engine) {
	movieListEndpoint(router)
	movieAddEndpoint(router)
	movieDetailsEndpoint(router)
	movieDeleteEndpoint(router)
}
