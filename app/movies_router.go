package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Find a better way that makes use of interfaces in the queryDatabase function to query the database.
// Since this would prevent a lot of code duplication.
func queryMovies(movies *[]*Movie) {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	rows, err := database.Query("SELECT IMDb_id, Title, Rating, Year FROM movies;")
	CheckError(err)
	defer closeRows(rows)

	for rows.Next() {
		var movie Movie
		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
		*movies = append(*movies, &movie)
	}

	defer rows.Close()
}

// TODO: IDEM, ideally the commands would all return structs that are serialized to JSON.
// But this requires a refactor, while I just want to get everything working for now.
func queryMovie(id *string) *Movie {
	database := OpenMoviesDatabase()
	defer CloseMoviesDatabase(database)

	sql := fmt.Sprintf(`
		SELECT IMDb_id, Title, Rating, Year
		FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	rows, err := database.Query(sql)
	CheckError(err)
	defer closeRows(rows)

	var movie Movie
	for rows.Next() {
		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
	}

	defer rows.Close()

	return &movie
}

func movieListEndpoint(router *gin.Engine) {
	router.GET("/movies", func(context *gin.Context) {
		var movies []*Movie
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
		DeleteMovie(&id)
		context.Status(http.StatusNoContent)
	})
}

func movieAddEndpoint(router *gin.Engine) {
	router.POST("/movies", func(context *gin.Context) {
		var movie Movie

		if err := context.BindJSON(&movie); err != nil {
			return
		}

		AddMovie(&movie)

		context.IndentedJSON(http.StatusOK, &movie)
	})
}

func AddApiEndpoints(router *gin.Engine) {
	movieListEndpoint(router)
	movieAddEndpoint(router)
	movieDetailsEndpoint(router)
	movieDeleteEndpoint(router)
}
