package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Find a better way that makes use of interfaces in the queryDatabase function to query the database.
// Since this would prevent a lot of code duplication.
func queryMovies(database *sql.DB, movies *[]*Movie) {
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
func queryMovie(database *sql.DB, id *string, movie *Movie) {
	sql := fmt.Sprintf(`
		SELECT IMDb_id, Title, Rating, Year
		FROM movies
		WHERE IMDb_id='%s';
	`, *id)

	rows, err := database.Query(sql)
	CheckError(err)
	defer closeRows(rows)

	for rows.Next() {
		rows.Scan(&movie.IMDbId, &movie.Title, &movie.IMDbRating, &movie.ReleaseYear)
	}

	defer rows.Close()
}

func movieListEndpoint(router *gin.Engine, database *sql.DB) {
	router.GET("/movies", func(context *gin.Context) {
		var movies []*Movie
		queryMovies(database, &movies)

		context.IndentedJSON(http.StatusOK, &movies)
	})
}

func movieDetailsEndpoint(router *gin.Engine, database *sql.DB) {
	router.GET("/movies/:id", func(context *gin.Context) {
		id := context.Param("id")
		var movie Movie
		queryMovie(database, &id, &movie)

		if movie.IMDbId != nil {
			context.IndentedJSON(http.StatusOK, &movie)
		} else {
			context.Status(404)
		}
	})
}

func movieDeleteEndpoint(router *gin.Engine, database *sql.DB) {
	router.DELETE("/movies/:id", func(context *gin.Context) {
		id := context.Param("id")
		DeleteMovie(database, &id)
		context.Status(http.StatusOK)
	})
}

func movieAddEndpoint(router *gin.Engine, database *sql.DB) {
	router.POST("/movies", func(context *gin.Context) {
		var movie Movie

		if err := context.BindJSON(&movie); err != nil {
			return
		}

		AddMovie(database, &movie)

		context.IndentedJSON(http.StatusOK, &movie)
	})
}

func AddApiEndpoints(router *gin.Engine, database *sql.DB) {
	movieListEndpoint(router, database)
	movieAddEndpoint(router, database)
	movieDetailsEndpoint(router, database)
	movieDeleteEndpoint(router, database)
}
